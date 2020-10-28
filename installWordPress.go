package main

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/mkideal/cli"
	"os"
)

var wordpressInstallCommand = &cli.Command{
	Name: "install-wp",
	Fn: func(ctx *cli.Context) error {
		installWordPress("")

		return nil
	},
}

func installWordPress(wordpressPath string) {

	if wordpressPath == "" {
		survey.AskOne(&survey.Input{Message: "Dans quel répertoire ?", Default: "wordpress"}, &wordpressPath)

		if wordpressPath == "" {
			os.Exit(1)
		}
	}

	println(wordpressPath)

	siteTitle := "Mon nouveau site"
	adminUser := "inrage"
	adminPassword := "123go"
	adminMail := "pascal@inrage.fr"

	cmd, _ := runCommand(
		fmt.Sprintf(
		"wp core install --path=%s --url=127.0.0.1:8000 --title='%s' --admin_user=%s --admin_password=%s --admin_email=%s --skip-email",
		wordpressPath, siteTitle, adminUser, adminPassword, adminMail,
		),
	)

	println(cmd)
}
