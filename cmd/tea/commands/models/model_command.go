package models

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jhontea/tea/cmd/tea/templates"
	"github.com/jhontea/tea/cmd/tea/utils"
)

type ModelCommand struct {
	ModelName   string
	TableName   string
	ObjectName  string
	PackageName string
}

func (command *ModelCommand) Help() {
	fmt.Printf(`
Description:
	The ./tea model command prints help about the given command.
	The command 'model' creates a gin API model.
Example:
	$ ./tea model [path] [model name] [table name]
`)
}

func (command *ModelCommand) Execute(args []string) {
	if len(args) < 3 {
		command.Help()
		os.Exit(2)
	}

	modelPath := args[0]
	modelName := args[1]
	tableName := args[2]

	command.ModelName = strings.Title(modelName)
	command.TableName = tableName
	command.ObjectName = "V1" + strcase.ToCamel(modelName)
	command.PackageName = strcase.ToSnake(modelName)

	log.Println("Creating model...")

	modelTemplate := "cmd/tea/templates/model.go.tmpl"
	content := templates.GenerateTemplate(modelTemplate, command)

	fileName := strings.ToLower(modelName)

	utils.WriteToPath(modelPath, fileName, content)

	log.Println("Creating object...")

	objectTemplate := "cmd/tea/templates/object.go.tmpl"
	content = templates.GenerateTemplate(objectTemplate, command)

	fileName = "v1_" + strings.ToLower(modelName) + "_object"
	objectPath := "src/objects/" + command.PackageName

	utils.WriteToPath(objectPath, fileName, content)
}
