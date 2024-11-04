package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var dirName string
var moduleName string
var capitalizedName string

type Variables struct {
	Name            string
	CapitalizedName string
	ModuleName      string
}

type Config struct {
	EntityTemplate     string `json:"entityTemplate"`
	InitTemplate       string `json:"initTemplate"`
	ServiceTemplate    string `json:"serviceTemplate"`
	ControllerTemplate string `json:"controllerTemplate"`
	RepositoryTemplate string `json:"repositoryTemplate"`
	BaseTemplate       string `json:"baseTemplate"`
	EntitiesDir        string `json:"entitiesDir"`
	RepositoryDir      string `json:"repositoryDir"`
	ServiceDir         string `json:"serviceDir"`
	InitDir            string `json:"initDir"`
	ControllerDir      string `json:"controllerDir"`
}

var config Config

func loadConfig() error {

	configPath := "/etc/codegen/config/config.json"
	if _, err := os.Stat(configPath); os.IsExist(err) {

	}
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Error opening config: %v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding config: %v", err)
	}
	return nil
}
func GenerateFiles(input string) error {
	if isCreatedBefore(input) {
		log.Fatalf("Error: files created before with given input : %s", input)

	}
	if err := loadConfig(); err != nil {

		log.Fatalf("Error: %v", err)
	}
	file, err := os.Open("go.mod")
	if err != nil {

		log.Fatalf("Error: %v", err)
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
	dirName = input
	capitalizedName = strings.ToUpper(dirName[:1]) + strings.ToLower(dirName[1:])
	variables := Variables{
		Name:            dirName,
		CapitalizedName: capitalizedName,
		ModuleName:      moduleName,
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	if err := generateController(variables); err != nil {
		log.Fatalf("Error: %v", err)
	}
	if err = generateEntity(variables); err != nil {
		log.Fatalf("Error: %v", err)
	}
	if err = generateInit(variables); err != nil {
		log.Fatalf("Error: %v", err)
	}
	if err = generateRepository(variables); err != nil {
		log.Fatalf("Error: %v", err)
	}
	if err = generateService(variables); err != nil {
		log.Fatalf("Error: %v", err)
	}
	return nil
}

// entity creation
func generateEntity(variables Variables) error {
	dirPath := config.EntitiesDir
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Fatalf("Error: %v", err)
	}

	entityTemplate := config.EntityTemplate
	tmpl, err := template.ParseFiles(entityTemplate)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("Creating etitiy class ...")
	entityFile, err := os.Create(fmt.Sprintf("%s/%s.go", dirPath, variables.CapitalizedName))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer entityFile.Close()
	log.Println("done")

	data := Variables{
		Name:            variables.Name,
		CapitalizedName: variables.CapitalizedName,
	}
	err = tmpl.Execute(entityFile, data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// base creation
	baseFilePath := fmt.Sprintf("%s/Base.go", dirPath)
	if _, err := os.Stat(baseFilePath); os.IsNotExist(err) {
		baseTemplate := config.BaseTemplate
		tmpl, err = template.ParseFiles(baseTemplate)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		log.Println("Creating base class ...")
		baseFile, err := os.Create(fmt.Sprintf("%s/Base.go", dirPath))

		if err != nil {
			log.Fatalf("Error creating Base.go: %v", err)
		}
		defer baseFile.Close()
		log.Println("done")

		err = tmpl.Execute(baseFile, nil)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		return nil
	}
	return nil
}

// repository creation
func generateRepository(variables Variables) error {
	dirPath := fmt.Sprintf(config.RepositoryDir+"/%s", dirName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Fatalf("Error: %v", err)
	}

	repositoryTemplate := config.RepositoryTemplate
	tmpl, err := template.ParseFiles(repositoryTemplate)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("Creating repository class ...")
	repositoryFile, err := os.Create(fmt.Sprintf("%s/%s/repository.go", config.RepositoryDir, dirName))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	data := Variables{
		Name:            variables.Name,
		CapitalizedName: variables.CapitalizedName,
		ModuleName:      variables.ModuleName,
	}
	err = tmpl.Execute(repositoryFile, data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer repositoryFile.Close()

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("done")

	return nil
}

// service creation
func generateService(variables Variables) error {
	dirPath := fmt.Sprintf(config.ServiceDir+"/%s", dirName)

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Fatalf("Error: %v", err)
	}
	serviceTemplate := config.ServiceTemplate
	tmpl, err := template.ParseFiles(serviceTemplate)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	data := Variables{
		Name:            variables.Name,
		CapitalizedName: variables.CapitalizedName,
		ModuleName:      variables.ModuleName,
	}
	log.Println("Creating service class ...")
	serviceFile, err := os.Create(fmt.Sprintf("%s/%s/service.go", config.ServiceDir, dirName))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer serviceFile.Close()
	err = tmpl.Execute(serviceFile, data)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("done")

	return nil
}

// controller creation
func generateController(variables Variables) error {
	fixedName := fixName(variables.Name)
	dirPath := fmt.Sprintf(config.ControllerDir+"/%s", fixedName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Fatalf("Error: %v", err)
	}
	controllerTemplate := config.ControllerTemplate

	tmpl, err := template.ParseFiles(controllerTemplate)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("Creating controller class ...")
	controllerFile, err := os.Create(fmt.Sprintf("%s/%s/%s.go", config.ControllerDir, fixedName, fixedName))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer controllerFile.Close()
	data := Variables{
		Name:            fixedName,
		CapitalizedName: variables.CapitalizedName,
		ModuleName:      variables.ModuleName,
	}
	err = tmpl.Execute(controllerFile, data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("done")

	return nil
}

// init creation
func generateInit(variables Variables) error {
	dirPath := fmt.Sprintf(config.InitDir+"/%s", dirName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Fatalf("Error: %v", err)
	}
	initTemplate := config.InitTemplate

	tmpl, err := template.ParseFiles(initTemplate)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("Creating init class ...")
	initFile, err := os.Create(fmt.Sprintf("%s/%s.go", dirPath, variables.Name))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer initFile.Close()

	data := Variables{
		Name:            variables.Name,
		CapitalizedName: variables.CapitalizedName,
		ModuleName:      variables.ModuleName,
	}
	err = tmpl.Execute(initFile, data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("done")

	return nil
}

func fixName(name string) string {
	nameChars := strings.Split(name, "")
	lastChar := nameChars[len(nameChars)-1]
	var fixed string
	switch lastChar {
	case "s":
		fixed = name + "es"
	case "y":
		fixed = name[:len(name)-1] + "ies"
	default:
		fixed = name + "s"
	}
	return fixed
}

func isCreatedBefore(input string) bool {
	basePath := "./app/services/"
	filePath := filepath.Join(basePath, fmt.Sprintf("%s/%s.go", input, input))
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else if err != nil {
		log.Fatalf("Error : %v", err)
	} else {
		return true
	}
	return false
}
func DeleteFiles(input string) {

	basePath := "./app/services/"

	filePath := filepath.Join(basePath, fmt.Sprintf("%s/%s.go", input, input))
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Dosya yok:", filePath)
	} else if err != nil {
		fmt.Println("Dosyaya erişim hatası:", err)
	} else {
		fmt.Println("Dosya mevcut:", filePath)
		deleteEntities(input)
		deleteControllers(input)
		deleteServices(input)
	}

}

func deleteControllers(input string) error {
	deleteFile := fixName(input)
	basePath := "./app/controllers"
	filePath := filepath.Join(basePath, fmt.Sprintf("%s", deleteFile))

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("Error: %v", err)
	} else if err != nil {
		log.Fatalf("Error: %v", err)
	} else {
		fmt.Println("controllers mevcut:", filePath)

		err := os.RemoveAll("./" + filePath)
		if err != nil {
			log.Fatalf("Error: %v", err)
		} else {
			fmt.Println("delete is success " + filePath)
			return nil
		}

	}
	return nil

}

func deleteEntities(input string) error {
	deleteFile := strings.ToUpper(input[:1]) + strings.ToLower(input[1:])
	fmt.Println(deleteFile)
	basePath := "./app/entities"
	filePath := filepath.Join(basePath, fmt.Sprintf("%s.go", deleteFile))

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("Error: %v", err)
	} else if err != nil {
		log.Fatalf("Error: %v", err)
	} else {
		fmt.Println("entities mevcut:", filePath)

		err := os.RemoveAll("./" + filePath)
		if err != nil {
			log.Fatalf("Error: %v", err)
		} else {
			fmt.Println("delete is success " + filePath)
			return nil
		}

	}
	return nil

}

func deleteServices(input string) error {
	deleteFile := input
	fmt.Println(deleteFile)
	basePath := "./app/services"
	filePath := filepath.Join(basePath, fmt.Sprintf("%s", deleteFile))

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("Error: %v", err)
	} else if err != nil {
		log.Fatalf("Error: %v", err)
	} else {
		fmt.Println("services mevcut:", filePath)

		err := os.RemoveAll("./" + filePath)
		if err != nil {
			log.Fatalf("Error: %v", err)
		} else {
			fmt.Println("delete is success " + filePath)
			return nil
		}

	}
	return nil

}
