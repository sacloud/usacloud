package main

import (
	"bytes"
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/schema"
	"github.com/sacloud/usacloud/tools"
	"gopkg.in/urfave/cli.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/template"
)

var (
	destination = "src/github.com/sacloud/usacloud/contrib/completion/bash/usacloud"
	ctx         = tools.NewGenerateContext()
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tgen-bash-completion\n")
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen-bash-completion: ")

	// sort by key
	resources := sortableResources{}
	for key, r := range ctx.ResourceDef {
		resources = append(resources, sortableResource{
			key:      key,
			resource: r,
		})
	}
	sort.Sort(resources)
	src, err := generateResource(resources)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}

	// Write to file.
	outputName := filepath.Join(ctx.Gopath(), destination)
	err = ioutil.WriteFile(outputName, []byte(src), 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
	fmt.Printf("generated: %s\n", destination)
}

type sortableResource struct {
	key      string
	resource *schema.Resource
}
type sortableResources []sortableResource

func (s sortableResources) Len() int {
	return len(s)
}

func (s sortableResources) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortableResources) Less(i, j int) bool {
	return s[i].key < s[j].key
}

func generateResource(resources sortableResources) (string, error) {

	commands := []string{}
	globalFlags := []string{}
	globalWantValues := []string{}

	// build commands
	for _, r := range resources {
		commands = append(commands, tools.ToDashedName(r.key))
		commands = append(commands, r.resource.Aliases...)
	}

	toFlagNames := func(names []string) []string {
		res := []string{}
		for _, n := range names {
			res = append(res, tools.ToCLIFlagName(n))
		}
		return res
	}

	for _, f := range command.GlobalFlags {
		switch f.(type) {
		case *cli.BoolFlag:
			globalFlags = append(globalFlags, toFlagNames(f.Names())...)
		default:
			globalWantValues = append(globalWantValues, toFlagNames(f.Names())...)

		}
	}

	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(srcTemplate))
	err := t.Execute(buf, map[string]interface{}{
		"Commands":         commands,
		"GlobalFlags":      globalFlags,
		"GlobalWantValues": globalWantValues,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

var srcTemplate = `
_usacloud() {
    COMPREPLY=();
    local commands=({{range .Commands}}{{.}} {{end}})

    local flags=({{range .GlobalFlags}}{{.}} {{end}}--help --version)
    local wants_value=({{range .GlobalWantValues}}{{.}} {{end}})
    local configflags=(--help --token --secret --zone --config --profile --show)
    local zones=(is1a is1b tk1a tk1v)

    local cur prev words cword
    _get_comp_words_by_ref -n : cur prev words cword
    local i
    local command=usacloud

    for (( i=1; i < ${cword}; ++i)); do
        local word=${words[i]}
        if [[ " ${wants_value[*]}  " =~ " ${word} " ]]; then
            # skip the next option
            (( ++i ))
        elif [[ " ${commands[*]} " =~ " ${word} " ]]; then
            command=${word}
            break
        fi
    done

    if [ "$command" == "usacloud" ]; then
        if [[ "$cur" == -* ]]; then
            COMPREPLY=($(compgen -W "${flags[*]} ${wants_value[*]}" -- "${cur}"))
        else
            if [[ " ${wants_value[*]}  " =~ " ${prev} " ]]; then
                if [ "--zone" == "${prev}" ]; then
                    COMPREPLY=($(compgen -W "${zones[*]}" -- "${cur}"))
                fi
                if [ "--profile" == "${prev}" -o "--config" == "${prev}" ]; then
                    COMPREPLY=($(compgen -W "$(${words[0]} config list)" -- "${cur}"))
                fi
            else
                COMPREPLY=($(compgen -W "${commands[*]}" -- "${cur}"))
            fi
        fi
    else
        if [ "$cur" = "${words[$i+1]}" ]; then
            opts=$( ${words[@]:0:$cword} --generate-completion );
            COMPREPLY=($(compgen -W "${opts}" -- "${cur}"))
        else
            opts=$( ${words[@]:0:$i+2} --generate-completion ${words[@]:$i+3} -- "${cur}" "${prev}" "${words[$i+1]}" );
            COMPREPLY=($(compgen -W "${opts}" -- "${cur}"))
        fi
    fi

    return 0
};


complete -F _usacloud -o default usacloud usacloud.exe
`
