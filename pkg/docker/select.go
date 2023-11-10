package docker

import (
	"docker-manager/pkg/utils" // Assuming utils is the package where PathExists is located
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// ANSI color codes for red and bold text
const (
	Red   = "\033[31m"
	Bold  = "\033[1m"
	Reset = "\033[0m"
)

func formatProjectName(name string, path string) string {
	if !utils.PathExists(path) {
		return fmt.Sprintf("%s%s%s %s(Path does not exist)%s", Bold, name, Reset, Red, Reset)
	}
	return fmt.Sprintf("%s%s%s", Bold, name, Reset)
}

func SelectProjects() ([]string, error) {
	projectNames := GetSortedProjectNames()
	formattedProjectNames := make([]string, len(projectNames))

	for i, name := range projectNames {
		projectPath := Projects[name]
		formattedProjectNames[i] = formatProjectName(name, projectPath)
	}

	var selectedProjects []string
	prompt := &survey.MultiSelect{
		Message: "Which projects do you want to start?",
		Options: formattedProjectNames,
	}
	err := survey.AskOne(prompt, &selectedProjects)
	if err != nil {
		return nil, err
	}

	for i, name := range selectedProjects {
		selectedProjects[i] = strings.Split(name, " ")[0]
	}

	return selectedProjects, nil
}
