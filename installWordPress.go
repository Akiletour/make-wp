package main

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/mkideal/cli"
	"os"
	"strings"
)

var wordpressInstallCommand = &cli.Command{
	Name: "install-wp",
	Fn: func(ctx *cli.Context) error {
		installWordPress()

		return nil
	},
}

func installWordPress() {

	if wordpressPath == "" {
		survey.AskOne(&survey.Input{Message: "Dans quel répertoire ?", Default: "wordpress"}, &wordpressPath)

		if wordpressPath == "" {
			os.Exit(1)
		}
	}

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

	postInstallWordpress()
}

func postInstallWordpress() {
	commander := fmt.Sprintf("wp --path=%s", wordpressPath)
	runCommand(fmt.Sprintf("%s post delete 1 2 --force", commander))
	runCommand(fmt.Sprintf("%s comment delete 1 --force", commander))
	runCommand(fmt.Sprintf("%s rewrite structure '/%%postname%%/' --hard", commander))
	runCommand(fmt.Sprintf("%s rewrite flush --hard", commander))

	installExtensions()
}

func installExtensions() {
	commander := fmt.Sprintf("wp --path=%s", wordpressPath)

	listPlugins := []string{
		"wordpress-seo",
		"contact-form-7",
		"classic-editor",
		"duplicate-post",
		"custom-post-type-ui",
	}

	sanitizePlugins := strings.Trim(fmt.Sprint(listPlugins), "[]")

	cmd, _ := runCommand(fmt.Sprintf("%s plugin install %v --activate --force", commander, sanitizePlugins))

	println(cmd)

	// Install ACF Pro with key directly from source
	acfProKey := ""

	survey.AskOne(&survey.Password{Message: "Clé ACF ?"}, &acfProKey)

	if acfProKey != "" {
		acfCmd, _ := runCommand(fmt.Sprintf("%s plugin install --activate \"http://connect.advancedcustomfields.com/index.php?p=pro&a=download&k=%s\"", commander, acfProKey))

		println(acfCmd)
	}
}
