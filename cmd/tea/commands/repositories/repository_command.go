package repositories

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jhontea/tea/cmd/tea/templates"
	"github.com/jhontea/tea/cmd/tea/utils"
)

type RepositoryCommand struct {
	ModelName      string
	RepositoryName string
	PackageName    string
}

func (command *RepositoryCommand) Help() {
	fmt.Printf(`
Description:
	The ./tea repository command prints help about the given command.
	The command 'repository' creates a gin API repository.
Example:
	$ ./tea repository [path] [repository name]
`)
}

func (command *RepositoryCommand) Execute(args []string) {
	if len(args) < 2 {
		command.Help()
		os.Exit(2)
	}

	repositoryPath := args[0]
	repositoryName := args[1]

	command.ModelName = strings.Title(repositoryName)
	command.RepositoryName = "V1" + strings.Title(repositoryName)
	command.PackageName = strcase.ToSnake(repositoryName)

	log.Println("Creating reppository...")

	modelTemplate := "cmd/tea/templates/repository.go.tmpl"
	content := templates.GenerateTemplate(modelTemplate, command)

	fileName := "v1_" + strcase.ToSnake(repositoryName) + "_repository"

	utils.WriteToPath(repositoryPath, fileName, content)
}
