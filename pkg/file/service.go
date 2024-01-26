package file

import (
	"os"
	"regexp"
)

func GetRootDirectory() string {
	projectDir := os.Getenv("PROJECT_DIR")
	projectName := regexp.MustCompile(`^(.*` + projectDir + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}