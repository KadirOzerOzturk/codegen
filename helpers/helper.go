package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dirName string
var moduleName string
var capitalizedName string

type Task struct {
	name    string
	execute func() error
}

func GenerateFiles(input string) error {
	tasks := []Task{
		{"Entity", genereteEntity},
		{"Controller", genereteController},
		{"Init", genereteInit},
		{"Repository", genereteRepository},
		{"Service", genereteService},
	}
	dirName = input
	capitalizedName = strings.ToUpper(dirName[:1]) + strings.ToLower(dirName[1:])

	file, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			moduleName = strings.TrimSpace(strings.TrimPrefix(line, "module "))
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	for _, task := range tasks {
		if err := task.execute(); err != nil {
			return fmt.Errorf("%s generation failed: %w", task.name, err)
		}
	}

	return nil
}

// entity creation
func genereteEntity() error {
	dirPath := "./app/entities"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}
	entityTemplate := "./templates/entityTemplate"

	entityContent, err := os.ReadFile(entityTemplate)
	if err != nil {
		panic(err)
	}
	entityTemplateStr := string(entityContent)
	entity := fmt.Sprintf(entityTemplateStr, capitalizedName)
	err = os.WriteFile(fmt.Sprintf("./app/entities/%s.go", dirName), []byte(entity), 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

// repository creation
func genereteRepository() error {
	repositoryTemplate := "./templates/repositoryTemplate"

	repoContent, err := os.ReadFile(repositoryTemplate)
	if err != nil {
		return err
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

	dirPath := fmt.Sprintf("./app/services/%s", dirName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("./app/services/%s/repository.go", dirName), []byte(repository), 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

// service creation
func genereteService() error {
	serviceTemplate := "./templates/serviceTemplate"

	serviceContent, err := os.ReadFile(serviceTemplate)
	if err != nil {
		return err
	}
	serviceTemplateStr := string(serviceContent)
	service := fmt.Sprintf(serviceTemplateStr,
		dirName,
		moduleName,
		moduleName,
		moduleName,
		capitalizedName,
		capitalizedName,
		capitalizedName,
		capitalizedName,
		capitalizedName,
		capitalizedName)

	err = os.WriteFile(fmt.Sprintf("./app/services/%s/service.go", dirName), []byte(service), 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

// controller creation
func genereteController() error {
	dirPath := fmt.Sprintf("./app/controllers/%ss", dirName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
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
		return err
	}
	return nil
}

// init creation
func genereteInit() error {
	dirPath := fmt.Sprintf("./app/services/%s", dirName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		err.Error()
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
		return err
	}
	return nil
}
