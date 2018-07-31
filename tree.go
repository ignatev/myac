package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

//└ ─ │ ├
const (
	newLine      = "\n"
	emptySpace   = "    "
	middleItem   = "├── "
	continueItem = "│   "
	lastItem     = "└── "
)

type configDirectory struct {
	currentDirPath string
	parentDirPath string
	filePaths      []string
	dirPaths       []string
}

func tree(currentDir, parentDir string) (string, error) {

	files, err := ioutil.ReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}

	cd := configDirectory{}
	for _, file := range files {
		if !file.IsDir() {
			cd.filePaths = append(cd.filePaths, file.Name())
		} 
		if file.IsDir() && file.Name() != ".git" { //todo add exclude group into config
			cd.dirPaths = append(cd.dirPaths, currentDir + "/" + file.Name())
		}
		cd.currentDirPath = currentDir
		cd.parentDirPath = parentDir
	}
	
	printConfigDirectory(cd)
	
	for _, dir := range cd.dirPaths {
		tree(dir, currentDir)
	}

	return "", nil

}

func printConfigDirectory(cd configDirectory) {
	fmt.Println("directories:")
	for _, dir := range cd.dirPaths {
		fmt.Println(dir)
	}
	fmt.Println("files:")
	for _, file := range cd.filePaths {
		fmt.Println(file)
	}
	fmt.Println("parentDir:", cd.parentDirPath)
	fmt.Println("currentDir:", cd.currentDirPath)
	fmt.Println("----")
	fmt.Println()
}
