// Copyright 2017-2025 The sacloud/usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// sdk-interface-inventory writes a stable JSON inventory of exported SDK interfaces.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"unicode"
)

type packageMeta struct {
	ImportPath string
	Dir        string
	GoFiles    []string
	CgoFiles   []string
	Export     string
}

type manifest struct {
	SchemaVersion int                 `json:"schema_version"`
	Package       string              `json:"package"`
	Interfaces    []interfaceManifest `json:"interfaces"`
}

type interfaceManifest struct {
	Interface string   `json:"interface"`
	Resource  string   `json:"resource"`
	Methods   []method `json:"methods"`
}

type method struct {
	Name      string   `json:"name"`
	Params    []string `json:"params"`
	Results   []string `json:"results"`
	Operation string   `json:"operation"`
	Command   string   `json:"command"`
}

func main() {
	var packagePath, interfaceName, output string
	flag.StringVar(&packagePath, "package", "", "go package import path")
	flag.StringVar(&interfaceName, "interface", "", "exported interface name (empty means all)")
	flag.StringVar(&output, "output", "-", "output JSON path, or - for stdout")
	flag.Parse()

	if packagePath == "" {
		exitf("--package is required")
	}
	meta, err := listPackage(packagePath)
	if err != nil {
		exitf("go list: %v", err)
	}
	inventory, err := inventoryPackage(meta, interfaceName)
	if err != nil {
		exitf("inventory: %v", err)
	}
	data, err := json.MarshalIndent(inventory, "", "  ")
	if err != nil {
		exitf("encode JSON: %v", err)
	}
	data = append(data, '\n')
	if output == "-" {
		_, err = os.Stdout.Write(data)
	} else {
		err = os.WriteFile(output, data, 0o644)
	}
	if err != nil {
		exitf("write output: %v", err)
	}
}

func exitf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "sdk-interface-inventory: "+format+"\n", args...)
	os.Exit(1)
}

func listPackage(path string) (packageMeta, error) {
	cmd := exec.Command("go", "list", "-json", "-export", path)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return packageMeta{}, fmt.Errorf("%w: %s", err, strings.TrimSpace(stderr.String()))
	}
	var meta packageMeta
	if err := json.Unmarshal(out, &meta); err != nil {
		return packageMeta{}, err
	}
	if meta.Dir == "" || meta.ImportPath == "" {
		return packageMeta{}, fmt.Errorf("incomplete go list result for %q", path)
	}
	return meta, nil
}

func inventoryPackage(meta packageMeta, requested string) (manifest, error) {
	fset := token.NewFileSet()
	files := make([]*ast.File, 0, len(meta.GoFiles)+len(meta.CgoFiles))
	for _, name := range append(append([]string{}, meta.GoFiles...), meta.CgoFiles...) {
		file, err := parser.ParseFile(fset, filepath.Join(meta.Dir, name), nil, parser.ParseComments)
		if err != nil {
			return manifest{}, err
		}
		files = append(files, file)
	}
	if len(files) == 0 {
		return manifest{}, fmt.Errorf("no Go files in %s", meta.Dir)
	}
	config := types.Config{Importer: moduleImporter(fset)}
	pkg, err := config.Check(meta.ImportPath, fset, files, nil)
	if err != nil {
		return manifest{}, fmt.Errorf("type check %s: %w", meta.ImportPath, err)
	}

	names := pkg.Scope().Names()
	sort.Strings(names)
	result := manifest{SchemaVersion: 2, Package: meta.ImportPath}
	found := false
	for _, name := range names {
		if requested != "" && name != requested {
			continue
		}
		if !ast.IsExported(name) {
			continue
		}
		named, ok := pkg.Scope().Lookup(name).Type().(*types.Named)
		if !ok {
			continue
		}
		iface, ok := named.Underlying().(*types.Interface)
		if !ok {
			continue
		}
		if requested == "" && !strings.HasSuffix(name, "API") {
			continue
		}
		found = true
		iface.Complete()
		entry := interfaceManifest{
			Interface: name,
			Resource:  resourceName(name),
		}
		for i := 0; i < iface.NumMethods(); i++ {
			entry.Methods = append(entry.Methods, newMethod(iface.Method(i)))
		}
		sort.Slice(entry.Methods, func(i, j int) bool { return entry.Methods[i].Name < entry.Methods[j].Name })
		result.Interfaces = append(result.Interfaces, entry)
	}
	if requested != "" && !found {
		return manifest{}, fmt.Errorf("exported interface %q not found", requested)
	}
	return result, nil
}

func newMethod(fn *types.Func) method {
	sig := fn.Type().(*types.Signature)
	m := method{Name: fn.Name(), Operation: classify(fn.Name())}
	m.Command = commandName(fn.Name(), m.Operation)
	for i := 0; i < sig.Params().Len(); i++ {
		m.Params = append(m.Params, typeString(sig.Params().At(i).Type()))
	}
	for i := 0; i < sig.Results().Len(); i++ {
		m.Results = append(m.Results, typeString(sig.Results().At(i).Type()))
	}
	return m
}

func typeString(t types.Type) string {
	return types.TypeString(t, func(p *types.Package) string {
		if p.Name() == "context" {
			return "context"
		}
		return p.Name()
	})
}

func classify(name string) string {
	switch {
	case strings.HasPrefix(name, "List"), strings.HasPrefix(name, "Find"):
		return "list"
	case strings.HasPrefix(name, "Read"), strings.HasPrefix(name, "Get"):
		return "read"
	case strings.HasPrefix(name, "Create"):
		return "create"
	case strings.HasPrefix(name, "Update"):
		return "update"
	case strings.HasPrefix(name, "Delete"):
		return "delete"
	default:
		return "action"
	}
}

func commandName(name, operation string) string {
	if operation != "action" {
		for _, prefix := range []string{"List", "Find", "Read", "Get", "Create", "Update", "Delete"} {
			if suffix := strings.TrimPrefix(name, prefix); suffix != name {
				if suffix == "" {
					return operation
				}
				return operation + "-" + kebabCase(suffix)
			}
		}
	}
	return kebabCase(name)
}

func kebabCase(name string) string {
	runes := []rune(name)
	var words strings.Builder
	for i, r := range runes {
		if i > 0 && unicode.IsUpper(r) &&
			(unicode.IsLower(runes[i-1]) || unicode.IsDigit(runes[i-1]) ||
				(unicode.IsUpper(runes[i-1]) && i+1 < len(runes) && unicode.IsLower(runes[i+1]))) {
			words.WriteByte('-')
		}
		words.WriteRune(unicode.ToLower(r))
	}
	return words.String()
}

func resourceName(interfaceName string) string {
	return kebabCase(strings.TrimSuffix(interfaceName, "API"))
}

// moduleImporter makes go/types module-aware without a dependency on x/tools.
// go list supplies the compiler export data that go/importer consumes.
func moduleImporter(fset *token.FileSet) types.Importer {
	var mu sync.Mutex
	exports := make(map[string]string)
	lookup := func(path string) (io.ReadCloser, error) {
		mu.Lock()
		export, ok := exports[path]
		mu.Unlock()
		if !ok {
			meta, err := listPackage(path)
			if err != nil {
				return nil, err
			}
			export = meta.Export
			if export == "" {
				return nil, fmt.Errorf("go list returned no export data for %q", path)
			}
			mu.Lock()
			exports[path] = export
			mu.Unlock()
		}
		return os.Open(export)
	}
	return importer.ForCompiler(fset, "gc", lookup)
}
