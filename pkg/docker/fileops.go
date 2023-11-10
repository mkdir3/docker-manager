package docker

import (
	"docker-manager/pkg/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func LoadProjectsFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Projects)
	if err != nil {
		return err
	}

	return nil
}

func SaveProjectsToFile(filename string) error {
	data, err := json.Marshal(Projects)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func CheckAndLoadProjectsFile(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		var createFile string
		prompt := &survey.Select{
			Message: fmt.Sprintf("The project configuration file '%s' does not exist. Would you like to create one and add projects? 😎", filePath),
			Options: []string{"Yes", "No"},
		}
		survey.AskOne(prompt, &createFile)
		if createFile == "Yes" {
			AddProject()
			SaveProjectsToFile(filePath)
		} else {
			utils.ExitInfo()
			os.Exit(0)
		}
	} else {
		if err := LoadProjectsFromFile(filePath); err != nil {
			fmt.Println()
			fmt.Printf("⚠️ Failed to load projects: %v\n", err)
			utils.ExitInfo()
		}
	}
	return nil
}
