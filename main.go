package main

import (
	"fmt"
	"os"

	"github.com/jhontea/tea/cmd/tea/commands"
)

func printUsageAndExit() {
	fmt.Printf("Go Scaffold like drinking a tea\n")
	for commandName, command := range commands.Commands {
		fmt.Printf("\nCommand `%s`:\n\n", commandName)
		command.Help()
	}
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		printUsageAndExit()
	}

	commandName := os.Args[1]
	commandArgs := []string{}

	if len(os.Args) > 2 {
		commandArgs = os.Args[2:]
	}

	command := commands.FindCommand(commandName)
	if command == nil {
		printUsageAndExit()
	}

	command.Execute(commandArgs)
}
