package main

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"io/ioutil"
	"os"
	"strings"
)

func installSageTheme() {

	directoryTheme := ""
	survey.AskOne(&survey.Input{Message: "Nom du répertoire du thème?", Default: "theme"}, &directoryTheme)

	if directoryTheme == "" {
		return
	}

	runCommandTty(fmt.Sprintf(
		"%s/wp-content/themes", wordpressPath),
		"composer", "create-project", "roots/sage", directoryTheme,
	)

	functionFile := fmt.Sprintf("%s/wp-content/themes/%s/resources/functions.php", wordpressPath, directoryTheme)

	read, _ := ioutil.ReadFile(functionFile)

	newContent := strings.Replace(string(read), "'setup', 'filters', 'admin'", "'setup', 'filters', 'admin', 'acf'", -1)
	ioutil.WriteFile(functionFile, []byte(newContent), 0)

	runCommandTty(fmt.Sprintf(
		"%s/wp-content/themes/%s", wordpressPath, directoryTheme),
		"composer", "require", "wordplate/acf",
	)

	os.Mkdir(
		fmt.Sprintf("%s/wp-content/themes/%s/app/Fields", wordpressPath, directoryTheme),
		0755,
	)

	copyFile("templates/sage/acf/acf.php", fmt.Sprintf("%s/wp-content/themes/%s/app/acf.php", wordpressPath, directoryTheme))

	runCommand(fmt.Sprintf("wp --path=%s theme activate %s/resources", wordpressPath, directoryTheme))
}