package shared

import (
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var projectDirName string = "list"

func GetRootPath() string {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, err := os.Getwd()
	if err != nil {
		log.Println("shared / GetRootPath / ", err)
		return ""
	}
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	return string(rootPath)
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	path := strings.Split(b, "/")
	path = path[:len(path)-2]
	return strings.Join(path, "/")
}
