package services

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jhontea/tea/cmd/tea/templates"
	"github.com/jhontea/tea/cmd/tea/utils"
)

type ServiceCommand struct {
	ServiceName string
	PackageName string
}

func (command *ServiceCommand) Help() {
	fmt.Printf(`
Description:
	The ./tea service command prints help about the given command.
	The command 'service' creates a gin API service.
Example:
	$ ./tea service [path] [service name]
`)
}

func (command *ServiceCommand) Execute(args []string) {
	if len(args) < 2 {
		command.Help()
		os.Exit(2)
	}

	path := args[0]
	serviceName := args[1]

	command.ServiceName = "V1" + strings.Title(serviceName)
	command.PackageName = strcase.ToSnake(serviceName)

	log.Println("Creating reppository...")

	template := "cmd/tea/templates/service.go.tmpl"
	content := templates.GenerateTemplate(template, command)

	fileName := "v1_" + strcase.ToSnake(serviceName) + "_service"

	utils.WriteToPath(path, fileName, content)
}
