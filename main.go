package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"os"
)

func main() {
	if err := cli.Root(rootCommand,
		cli.Tree(installCommand),
		cli.Tree(wordpressInstallCommand),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
