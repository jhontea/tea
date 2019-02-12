package commands

import (
	"github.com/jhontea/tea/cmd/tea/commands/controllers"
	"github.com/jhontea/tea/cmd/tea/commands/models"
	"github.com/jhontea/tea/cmd/tea/commands/repositories"
	"github.com/jhontea/tea/cmd/tea/commands/services"
)

var (
	Commands = map[string]Base{
		"controller": &controllers.ControllerCommand{},
		"model":      &models.ModelCommand{},
		"repository": &repositories.RepositoryCommand{},
		"service":    &services.ServiceCommand{},
	}
)

func FindCommand(name string) Base {
	switch name {
	case "c":
		{
			name = "controller"
		}
	case "m":
		{
			name = "models"
		}
	case "r":
		{
			name = "repository"
		}
	case "s":
		{
			name = "service"
		}
	}

	return Commands[name]
}
