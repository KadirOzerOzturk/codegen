package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var dirName string
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
		dirName = os.Args[1]
	} else {
		fmt.Println("Run main program")
	}
	file, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var moduleName string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			moduleName = strings.TrimSpace(strings.TrimPrefix(line, "module "))
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	dirPath := fmt.Sprintf("./app/entities/%s", dirName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return
	}
	entityTemplate := "./templates/entityTemplate"

	entityContent, err := os.ReadFile(entityTemplate)
	if err != nil {
		panic(err)
	}
	entityTemplateStr := string(entityContent)
	capitalizedName := strings.ToUpper(dirName[:1]) + strings.ToLower(dirName[1:])
	entity := fmt.Sprintf(entityTemplateStr, capitalizedName)
	err = os.WriteFile(fmt.Sprintf("./app/entities/%s.go", dirName), []byte(entity), 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	repositoryTemplate := "./templates/repositoryTemplate"

	repoContent, err := os.ReadFile(repositoryTemplate)
	if err != nil {
		panic(err)
	}

	repositoryTemplateStr := string(repoContent)
	repository := fmt.Sprintf(
		repositoryTemplateStr,
		dirName,
		moduleName,
		moduleName,
		moduleName,
		capitalizedName,
		capitalizedName,
		capitalizedName,
		capitalizedName,
		capitalizedName,
		capitalizedName,
		capitalizedName,
		capitalizedName,
		capitalizedName,
		capitalizedName)

	dirPath = fmt.Sprintf("./app/services/%s", dirName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return
	}
	err = os.WriteFile(fmt.Sprintf("./app/services/%s/repository.go", dirName), []byte(repository), 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	//service creation
	serviceTemplate := "./templates/serviceTemplate"

	serviceContent, err := os.ReadFile(serviceTemplate)
	if err != nil {
		panic(err)
	}
	serviceTemplateStr := string(serviceContent)
	service := fmt.Sprintf(serviceTemplateStr, dirName, moduleName, moduleName, moduleName, capitalizedName, capitalizedName, capitalizedName, capitalizedName, capitalizedName, capitalizedName)
	err = os.WriteFile(fmt.Sprintf("./app/services/%s/service.go", dirName), []byte(service), 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	// init creation
	dirPath = fmt.Sprintf("./app/services/%s", dirName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return
	}
	initTemplate := "./templates/initTemplate"

	initContent, err := os.ReadFile(initTemplate)
	if err != nil {
		panic(err)
	}
	initTemplateStr := string(initContent)
	init := fmt.Sprintf(initTemplateStr, dirName, moduleName, moduleName, moduleName)
	err = os.WriteFile(fmt.Sprintf("./app/services/%s/init.go", dirName), []byte(init), 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	// controller creation
	dirPath = fmt.Sprintf("./app/controllers/%ss", dirName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return
	}
	controllerTemplate := "./templates/controllerTemplate"

	controllerContent, err := os.ReadFile(controllerTemplate)
	if err != nil {
		panic(err)
	}
	controllerTemplateStr := string(controllerContent)
	controller := fmt.Sprintf(controllerTemplateStr,
		dirName,
		moduleName,
		moduleName,
		dirName,
		moduleName,
		moduleName,
		dirName,
		capitalizedName,
		dirName,
		dirName,
		dirName,
		capitalizedName,
		dirName)

	err = os.WriteFile(fmt.Sprintf("./app/controllers/%ss/%ss.go", dirName, dirName), []byte(controller), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

}
