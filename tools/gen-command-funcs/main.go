package main

import (
	"bytes"
	"fmt"
	"github.com/sacloud/usacloud/schema"
	"github.com/sacloud/usacloud/tools"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var (
	destination = "src/github.com/sacloud/usacloud/command"
	ctx         = tools.NewGenerateContext()
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tgen-command-funcs\n")
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen-command-funcs: ")

	for k, r := range ctx.ResourceDef {
		ctx.R = k
		err := generateResource(r)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	}
}

func generateResource(resource *schema.Resource) error {

	// build commands
	for k, c := range resource.Commands {
		ctx.C = k
		src, err := generateCommands(c)

		if err != nil {
			return err
		}

		// Write to file.
		// like 'switch_command_list.go'
		baseName := ctx.CommandFileName(c.UseCustomCommand)
		outputName := filepath.Join(ctx.Gopath(), destination, baseName)

		// target file is exist?
		_, err = os.Stat(outputName)
		if !c.UseCustomCommand || err != nil {
			err = ioutil.WriteFile(outputName, tools.Sformat([]byte(src)), 0644)
			if err != nil {
				return err
			}
			fmt.Printf("generated: %s\n", filepath.Join(destination, baseName))
		}
	}

	return nil
}

func generateCommands(command *schema.Command) (string, error) {
	b := bytes.NewBufferString("")
	t := template.New("c")
	template.Must(t.Parse(commandTemplate))

	action, err := generateAction(command)
	if err != nil {
		return "", err
	}

	err = t.Execute(b, map[string]interface{}{
		"FuncName":  ctx.CommandFuncName(),
		"ParamName": ctx.InputModelTypeName(),
		"Action":    action,
	})

	return b.String(), err
}

var commandTemplate = `package command

import (
    "fmt"
)

func {{.FuncName}}(ctx Context, params *{{.ParamName}}) error {
    {{.Action}}
}
`

func generateAction(command *schema.Command) (string, error) {

	var res string
	var err error
	switch command.Type {
	case schema.CommandList:
		// list / find / search
		res, err = generateFindCommand(command)
	case schema.CommandCreate:
		// create
		res, err = generateCreateCommand(command)
	case schema.CommandRead:
		// read
		res, err = generateReadCommand(command)
	case schema.CommandUpdate:
		// update
		res, err = generateUpdateCommand(command)
	case schema.CommandDelete:
		// delete
		res, err = generateDeleteCommand(command)
	case schema.CommandManipulate:
		// power-on/off
		res, err = generateManipulateCommand(command)
	case schema.CommandCustom:
		// custom
		res = `return fmt.Errorf("Not implements")`
	}

	return res, err
}
