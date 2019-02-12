package commands

type Base interface {
	Execute(args []string)
	Help()
}
