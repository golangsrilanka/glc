package prompt

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/GolangSriLanka/glc/config"
	"github.com/gookit/color"
	"log"
)

func SelectProviderWithPrompt() {
	prompt := &survey.Select{
		Message: "Select Git Provider",
		Options: []string{"github", "gitlab", "bitbucket"},
	}

	var selectedProvider string
	if err := survey.AskOne(prompt, &selectedProvider); err != nil {
		log.Fatalf("Unable to select git provider %s: ", err)
		return
	}

	if selectedProvider == "" {
		yellow := color.Yellow.Render
		log.Fatalf(yellow("No organization has been selected"))
		return
	}

	config.SetGitProvider(selectedProvider)

	green := color.Green.Render
	log.Printf("%s has selected has git provider", green(selectedProvider))

}
