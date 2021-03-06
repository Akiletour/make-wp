package main

import (
	"fmt"
	"github.com/mkideal/cli"
)

type appT struct {
	Name    string
	Version string
}

var app = appT{
	"make-wp",
	"v1",
}

type rootT struct {
	cli.Helper
}

func printHeader() {
	fmt.Print(sprintHeader() + "\n\n")
}

var rootCommand = &cli.Command{
	Desc: sprintHeader(),
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		printHeader()

		fmt.Print("Usage: make-wp <command>\n")
		fmt.Printf("More information on usage with %s command.\n", "help")

		return nil
	},
}

func sprintHeader() string {
	return fmt.Sprintf("%s version %s", app.Name, app.Version)
}
