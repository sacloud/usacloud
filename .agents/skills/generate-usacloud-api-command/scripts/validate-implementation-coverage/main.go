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

// validate-implementation-coverage verifies manifest methods have explicit core.Command Func values.
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
	"strconv"
	"strings"
	"sync"
)

const corePackagePath = "github.com/sacloud/usacloud/pkg/core"

type manifest struct {
	Interfaces []interfaceManifest `json:"interfaces"`
}

type interfaceManifest struct {
	Interface string     `json:"interface"`
	Resource  string     `json:"resource"`
	Methods   []method   `json:"methods"`
	Excluded  []excluded `json:"excluded"`
}

type method struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

type excluded struct {
	Method string `json:"method"`
	Reason string `json:"reason"`
}

type packageMeta struct {
	ImportPath string
	Dir        string
	GoFiles    []string
	Export     string
}

func main() {
	var manifestPath, commandDir string
	var strict bool
	flag.StringVar(&manifestPath, "manifest", "", "inventory JSON manifest")
	flag.StringVar(&commandDir, "command-dir", "", "directory containing hand-written commands")
	flag.BoolVar(&strict, "strict", false, "also reject manifest methods without command names")
	flag.Parse()
	if manifestPath == "" || commandDir == "" {
		exitf("--manifest and --command-dir are required")
	}
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		exitf("read manifest: %v", err)
	}
	var m manifest
	if err := json.Unmarshal(data, &m); err != nil {
		exitf("decode manifest: %v", err)
	}
	if err := validateManifest(m, strict); err != nil {
		exitf("manifest: %v", err)
	}
	if err := validateCoverage(m, commandDir, strict); err != nil {
		exitf("%v", err)
	}
	fmt.Printf("validated %d SDK methods in %s\n", methodCount(m), commandDir)
}

func validateCoverage(m manifest, commandDir string, strict bool) error {
	commands, err := commandNames(commandDir)
	if err != nil {
		return fmt.Errorf("scan commands: %w", err)
	}
	var missing []string
	for _, api := range m.Interfaces {
		excluded := excludedMethods(api.Excluded)
		for _, method := range api.Methods {
			if excluded[method.Name] {
				continue
			}
			identity := commandIdentity{Resource: api.Resource, Command: method.Command}
			if !commands[identity] {
				missing = append(missing, fmt.Sprintf("%s/%s: %s -> %s (core.Command.Func)",
					api.Interface, api.Resource, method.Name, method.Command))
			}
		}
	}
	if len(missing) > 0 {
		sort.Strings(missing)
		return fmt.Errorf("uncovered SDK methods:\n  %s", strings.Join(missing, "\n  "))
	}
	return nil
}

func methodCount(m manifest) int {
	var count int
	for _, api := range m.Interfaces {
		excluded := excludedMethods(api.Excluded)
		for _, method := range api.Methods {
			if !excluded[method.Name] {
				count++
			}
		}
	}
	return count
}

func exitf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "validate-implementation-coverage: "+format+"\n", args...)
	os.Exit(1)
}

func validateManifest(m manifest, strict bool) error {
	if len(m.Interfaces) == 0 {
		return fmt.Errorf("no interfaces in manifest")
	}
	for _, api := range m.Interfaces {
		if api.Interface == "" || api.Resource == "" {
			return fmt.Errorf("every interface needs interface and resource")
		}
		knownMethods := make(map[string]bool, len(api.Methods))
		for _, method := range api.Methods {
			if method.Name == "" {
				return fmt.Errorf("%s: method name is empty", api.Interface)
			}
			knownMethods[method.Name] = true
		}
		for _, item := range api.Excluded {
			if item.Method == "" || strings.TrimSpace(item.Reason) == "" {
				return fmt.Errorf("%s: every excluded entry needs method and reason", api.Interface)
			}
			if !knownMethods[item.Method] {
				return fmt.Errorf("%s: excluded method %q is not inventoried", api.Interface, item.Method)
			}
		}
		for _, method := range api.Methods {
			if strict && method.Command == "" {
				return fmt.Errorf("%s/%s: method %q has no reviewed command name",
					api.Interface, api.Resource, method.Name)
			}
		}
	}
	return nil
}

func excludedMethods(excluded []excluded) map[string]bool {
	result := make(map[string]bool, len(excluded))
	for _, item := range excluded {
		result[item.Method] = true
	}
	return result
}

type commandIdentity struct {
	Resource string
	Command  string
}

func commandNames(dir string) (map[commandIdentity]bool, error) {
	metas, err := listPackages(dir)
	if err != nil {
		return nil, err
	}
	result := make(map[commandIdentity]bool)
	for _, meta := range metas {
		resource, commands, isResource, err := packageCommands(meta)
		if err != nil {
			return nil, err
		}
		if !isResource {
			continue
		}
		for name := range commands {
			result[commandIdentity{Resource: resource, Command: name}] = true
		}
	}
	return result, nil
}

func packageCommands(meta packageMeta) (string, map[string]bool, bool, error) {
	fset := token.NewFileSet()
	files := make([]*ast.File, 0, len(meta.GoFiles))
	for _, name := range meta.GoFiles {
		file, err := parser.ParseFile(fset, filepath.Join(meta.Dir, name), nil, 0)
		if err != nil {
			return "", nil, false, err
		}
		files = append(files, file)
	}
	if len(files) == 0 {
		return "", nil, false, fmt.Errorf("no non-test Go files in %s", meta.Dir)
	}
	info := types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	config := types.Config{Importer: moduleImporter(fset)}
	if _, err := config.Check(meta.ImportPath, fset, files, &info); err != nil {
		return "", nil, false, fmt.Errorf("type check: %w", err)
	}

	var resourceNames []string
	for _, file := range files {
		ast.Inspect(file, func(node ast.Node) bool {
			lit, ok := node.(*ast.CompositeLit)
			if !ok || !isCoreType(lit.Type, "Resource", info) {
				return true
			}
			if name := namedField(lit, "Name"); name != "" {
				resourceNames = append(resourceNames, name)
			}
			return true
		})
	}
	if len(resourceNames) == 0 {
		return "", nil, false, nil
	}
	sort.Strings(resourceNames)
	if len(resourceNames) != 1 {
		return "", nil, false, fmt.Errorf("expected one core.Resource name in %s, found %q", meta.Dir, resourceNames)
	}

	result := make(map[string]bool)
	for _, file := range files {
		ast.Inspect(file, func(node ast.Node) bool {
			lit, ok := node.(*ast.CompositeLit)
			if !ok || !isCoreType(lit.Type, "Command", info) {
				return true
			}
			name, explicitFunc := commandFields(lit)
			if name != "" && explicitFunc {
				result[name] = true
			}
			return true
		})
	}
	return resourceNames[0], result, true, nil
}

func listPackage(path string) (packageMeta, error) {
	if !filepath.IsAbs(path) && !strings.HasPrefix(path, ".") {
		if _, err := os.Stat(path); err == nil {
			path = "./" + path
		}
	}
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

func listPackages(dir string) ([]packageMeta, error) {
	if !filepath.IsAbs(dir) && !strings.HasPrefix(dir, ".") {
		if _, err := os.Stat(dir); err == nil {
			dir = "./" + dir
		}
	}
	pattern := strings.TrimSuffix(dir, "/") + "/..."
	cmd := exec.Command("go", "list", "-json", "-export", pattern)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err, strings.TrimSpace(stderr.String()))
	}
	decoder := json.NewDecoder(bytes.NewReader(out))
	var packages []packageMeta
	for {
		var meta packageMeta
		if err := decoder.Decode(&meta); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if meta.Dir == "" || meta.ImportPath == "" {
			return nil, fmt.Errorf("incomplete go list result for %q", dir)
		}
		packages = append(packages, meta)
	}
	if len(packages) == 0 {
		return nil, fmt.Errorf("no command packages found below %s", dir)
	}
	return packages, nil
}

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

func isCoreType(expr ast.Expr, name string, info types.Info) bool {
	t, ok := info.Types[expr].Type.(*types.Named)
	if !ok {
		return false
	}
	obj := t.Obj()
	return obj.Name() == name && obj.Pkg() != nil && obj.Pkg().Path() == corePackagePath
}

func commandFields(lit *ast.CompositeLit) (string, bool) {
	name := namedField(lit, "Name")
	var explicitFunc bool
	for _, elt := range lit.Elts {
		field, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		key, ok := field.Key.(*ast.Ident)
		if !ok {
			continue
		}
		switch key.Name {
		case "Func":
			if ident, ok := field.Value.(*ast.Ident); !ok || ident.Name != "nil" {
				explicitFunc = true
			}
		}
	}
	return name, explicitFunc
}

func namedField(lit *ast.CompositeLit, fieldName string) string {
	for _, elt := range lit.Elts {
		field, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		key, ok := field.Key.(*ast.Ident)
		if !ok || key.Name != fieldName {
			continue
		}
		value, ok := field.Value.(*ast.BasicLit)
		if !ok || value.Kind != token.STRING {
			return ""
		}
		decoded, err := strconvUnquote(value.Value)
		if err != nil {
			return ""
		}
		return decoded
	}
	return ""
}

func strconvUnquote(s string) (string, error) {
	return strconv.Unquote(s)
}
