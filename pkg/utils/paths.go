package utils

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func PathExists(path string) bool {
	loader := NewLoader()
	loader.Start("Resolving paths from projects.json... ")

	resolvedPath, err := ResolveHomeDir(path)
	if err != nil {
		loader.Stop()
		fmt.Printf("Error resolving path: %v\n", err)
		return false
	}

	_, err = os.Stat(resolvedPath)
	loader.Stop()
	return !os.IsNotExist(err)
}

func PathChecker(path string) {
	if !PathExists(path) {
		fmt.Printf("The path %s does not exist or is not accessible.\n", path)
		return
	}
}

func ResolveHomeDir(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return strings.Replace(path, "~", usr.HomeDir, 1), nil
}

func GetComposeFilePath(projectDir string) (string, error) {
	fileNames := []string{"docker-compose.yml", "docker-compose.yaml"}

	for _, fileName := range fileNames {
		fullPath := projectDir + "/" + fileName
		if PathExists(fullPath) {
			return fullPath, nil
		}
	}

	return "", fmt.Errorf("no docker-compose file found in %s 🤔", projectDir)
}
