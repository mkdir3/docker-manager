package docker

import (
	"docker-manager/pkg/utils"
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
)

var Projects = make(map[string]string)

func init() {
	if err := LoadProjectsFromFile("projects.json"); err != nil {
		Projects = make(map[string]string)
	}
}

func GetSortedProjectNames() []string {
	projectNames := make([]string, 0, len(Projects))
	for projectName := range Projects {
		projectNames = append(projectNames, projectName)
	}
	sort.Strings(projectNames)
	return projectNames
}

func AddProject() {
	var projectName, projectPath string
	survey.AskOne(&survey.Input{Message: "Enter the project name:"}, &projectName)
	survey.AskOne(&survey.Input{Message: "Enter the project path:"}, &projectPath)
	utils.PathChecker(projectPath)

	var confirm string
	confirmPrompt := &survey.Select{
		Message: fmt.Sprintf("Do you want to add project %s with path %s?", projectName, projectPath),
		Options: []string{"Yes", "No"},
	}
	survey.AskOne(confirmPrompt, &confirm)

	if confirm == "Yes" {
		Projects[projectName] = projectPath
		SaveProjectsToFile("projects.json")
		fmt.Printf("Added project %s\n", projectName)
	}
}

func RemoveProject() error {
	var projectName string
	survey.AskOne(&survey.Input{Message: "Enter the project name you'd like to remove:"}, &projectName)

	if _, exists := Projects[projectName]; !exists {
		return fmt.Errorf("project %s does not exist", projectName)
	}

	delete(Projects, projectName)

	if err := SaveProjectsToFile("projects.json"); err != nil {
		return fmt.Errorf("failed to save projects after removal: %v", err)
	}

	return nil
}

func RemoveSingleProject(projectName string) error {
	if _, exists := Projects[projectName]; !exists {
		return fmt.Errorf("project %s does not exist", projectName)
	}

	var confirm string
	confirmPrompt := &survey.Select{
		Message: fmt.Sprintf("Do you want to remove unexisting project %s from the list?", projectName),
		Options: []string{"Yes", "No"},
	}
	survey.AskOne(confirmPrompt, &confirm)

	if confirm == "Yes" {
		delete(Projects, projectName)
		if err := SaveProjectsToFile("projects.json"); err != nil {
			return fmt.Errorf("failed to save projects after removal: %v", err)
		}
		fmt.Printf("\nProject %s removed successfully.", projectName)
		return nil
	}

	fmt.Printf("\nProject %s has not been removed.", projectName)
	return nil
}
