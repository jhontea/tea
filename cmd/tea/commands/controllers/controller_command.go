package controllers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jhontea/tea/cmd/tea/templates"
	"github.com/jhontea/tea/cmd/tea/utils"
)

type ControllerCommand struct {
	ControllerName string
}

func (command *ControllerCommand) Help() {
	fmt.Printf(`
Description:
	The ./tea controller command prints help about the given command.
	The command 'controller' creates a gin API controller.
Example:
	$ ./tea controller [path] [controller name]
`)
}

func (command *ControllerCommand) Execute(args []string) {
	if len(args) < 2 {
		command.Help()
		os.Exit(2)
	}

	controllerPath := args[0]
	controllerName := args[1]

	command.ControllerName = "V1" + strings.Title(controllerName)

	log.Println("Creating controller...")

	controllerTemplate := "cmd/tea/templates/controller.go.tmpl"
	content := templates.GenerateTemplate(controllerTemplate, command)

	fileName := "v1_" + strings.ToLower(controllerName) + "_controller"

	utils.WriteToPath(controllerPath, fileName, content)
}
