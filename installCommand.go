package main

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/mkideal/cli"
	"log"
	"os"
)

var wordpressPath string

var installCommand = &cli.Command{
	Name: "install",
	Desc: "install needed binaries",
	Fn: func(ctx *cli.Context) error {
		fmt.Printf("Salut")

		exists, _ := checkBinaryExist("wp")

		if exists {

			installWp := false
			survey.AskOne(&survey.Confirm{Message: "Voulez-vous installer WordPress ?"}, &installWp)

			if installWp == false {
				os.Exit(1)
			}

			survey.AskOne(&survey.Input{Message: "Dans quel répertoire ?", Default: "wordpress"}, &wordpressPath)

			if wordpressPath == "" {
				os.Exit(1)
			}

			downloadWordPress()

			cleanUpDefaultWordPress()

			prepareWpConfig()
			
			createWpCliConfig()

			installWordPress()

			installSage := false
			survey.AskOne(&survey.Confirm{Message: "Est-ce que l'on doit installer le thème Sage ?"}, &installSage)
			if installSage == true {
				installSageTheme()
			}
		}

		return nil
	},
}

func createWpCliConfig() {
	copyFile("templates/wp-cli.yml", fmt.Sprintf("%s/wp-cli.yml", wordpressPath))
}

func downloadWordPress() {
	cmd, err := runCommand(fmt.Sprintf("wp core download --locale=fr_FR --path=%s", wordpressPath))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(cmd)
}

func cleanUpDefaultWordPress() {
	fileToRemove := []string{
		"wp-content/plugins/hello.php",
		"wp-content/plugins/akismet",
		"wp-content/themes/twentytwenty",
		"wp-content/themes/twentynineteen",
		"wp-content/themes/twentyseventeen",
	}

	for _, file := range fileToRemove {
		os.RemoveAll(fmt.Sprintf("%s/%s", wordpressPath, file))
	}
}

func prepareWpConfig() {
	var qs = []*survey.Question{
		{
			Name: "dbname",
			Prompt: &survey.Input{Message: "Nom de la base de donnée", Default: "wordpress"},
		},
		{
			Name: "dbuser",
			Prompt: &survey.Input{Message: "Utilisateur BDD", Default: "root"},
		},
		{
			Name: "dbpass",
			Prompt: &survey.Input{Message: "Mot de passe BDD", Default: "root"},
		},
	}

	answers := struct {
		DBName	string
		DBUser 	string
		DBPass 	string
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	runCommand(
		fmt.Sprintf(
			"wp config create --dbname=%s --dbuser=%s --dbpass=%s --locale=fr_FR --path=%s",
			answers.DBName, answers.DBUser, answers.DBPass, wordpressPath,
			),
		)

	cmd, err := runCommand(fmt.Sprintf("wp db create --path=%s", wordpressPath))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(cmd)
}
