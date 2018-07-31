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
	currentDirPath	string
	parentDir		*configDirectory
	filePaths		[]string
	dirPaths		[]configDirectory
}

func tree(currentDir configDirectory, parentDir string) (string, error) {
	files, err := ioutil.ReadDir(currentDir.currentDirPath)
	if err != nil {
		log.Fatal(err)
	}

	cd := configDirectory{}
	cd.currentDirPath = currentDir.currentDirPath

	for _, file := range files {
		if !file.IsDir() {
			cd.filePaths = append(cd.filePaths, file.Name())
		} 
		if file.IsDir() && file.Name() != ".git" { //todo add exclude group into config
			innercd := configDirectory{}
			innercd.currentDirPath = currentDir.currentDirPath + "/" + file.Name()
			innercd.parentDir = &cd
			cd.dirPaths = append(cd.dirPaths, innercd)
		}
	}
	printConfigDirectory(cd)
	for _, dir := range cd.dirPaths {
		tree(dir, currentDir.currentDirPath)
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
	fmt.Println("parentDir:", cd.parentDir)
	fmt.Println("currentDir:", cd.currentDirPath)
	fmt.Println("----")
	fmt.Println()
}
