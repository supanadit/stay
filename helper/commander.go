// Original Author Supan Adit Pratama

package helper

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Commander struct {
}

type GoEnvironment struct {
	name  string
	value string
}

type Package struct {
	name     string
	host     string
	path     string
	fullPath string
}

var sourceName string = "src"

func (Commander) GetEnvironment() []GoEnvironment {
	command := "env"
	cmd := exec.Command("go", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Make sure golang have been installed")
	}
	outputData := string(out)
	outputDataSplit := strings.Split(outputData, "\n")
	var goEnvList []GoEnvironment
	for _, data := range outputDataSplit {
		if data != "" {
			commandValueSplit := strings.Split(data, "=")
			var goEnvData = GoEnvironment{
				name:  commandValueSplit[0],
				value: commandValueSplit[1],
			}
			lengthValue := len(goEnvData.value)
			if lengthValue != 0 {
				value := goEnvData.value
				var temporaryValue []string
				arrayString := strings.Split(value, "")
				for _, data := range arrayString {
					if data != "\"" {
						temporaryValue = append(temporaryValue, data)
					}
				}
				goEnvData.value = strings.Join(temporaryValue, "")
			}
			goEnvList = append(goEnvList, goEnvData)
		}
	}
	return goEnvList
}

func (Commander) GoPathSourcePackage() []Package {
	environment := GoEnvironment{}
	var location = environment.GoPath().value + "/" + sourceName
	var listPackageVendorHost []string // Eg. github.com, gopkg.in, golang.org
	files, err := ioutil.ReadDir(location)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if len(strings.Split(f.Name(), ".")) != 1 {
			listPackageVendorHost = append(listPackageVendorHost, location+"/"+f.Name())
		}
	}
	var listPackage []Package
	for _, data := range listPackageVendorHost {
		files, err := ioutil.ReadDir(data)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			splitPath := strings.Split(data, "/")
			hostName := splitPath[len(splitPath)-1]
			packageData := Package{}
			packageData.name = f.Name()
			packageData.host = hostName
			packageData.path = hostName + "/" + f.Name()
			packageData.fullPath = data + "/" + f.Name()
			listPackage = append(listPackage, packageData)
		}
	}
	return listPackage
}

func (Commander) Get(packageName string) {
	environment := GoEnvironment{}
	sourceGoPath := environment.GoPath().value + "/" + sourceName
	location := sourceGoPath + "/" + packageName
	temporaryDirectoryPath := sourceGoPath
	for _, directory := range strings.Split(packageName, "/") {
		temporaryDirectoryPath = temporaryDirectoryPath + "/" + directory
		stat, err := os.Stat(temporaryDirectoryPath)
		if !(err == nil && stat.IsDir()) {
			_ = os.MkdirAll(temporaryDirectoryPath, os.ModePerm)
		}
	}
	fullURL := "https://" + packageName
	_, err := git.PlainClone(location, false, &git.CloneOptions{
		URL:      fullURL,
		Progress: os.Stdout,
	})
	if err != nil {
		if err.Error() == "repository already exists" {
			fmt.Println("Already Downloaded")
		} else {
			fmt.Println(err)
		}
	}
}

func (GoEnvironment) GoEnv(environment string) GoEnvironment {
	goPathName := environment
	commander := Commander{}
	var environmentResult GoEnvironment
	for _, environment := range commander.GetEnvironment() {
		if environment.name == goPathName {
			environmentResult = environment
		}
	}
	return environmentResult
}

func (GoEnvironment) GoPath() GoEnvironment {
	environment := GoEnvironment{}
	return environment.GoEnv("GOPATH")
}
