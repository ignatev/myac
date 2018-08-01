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
	parentDir      *configDirectory
	filePaths      []string
	dirs           []configDirectory
}

func tree(currentDirPath string, parentDir *configDirectory) configDirectory {
	files, err := ioutil.ReadDir(currentDirPath)
	if err != nil {
		log.Fatal(err)
	}

	filePrefix := currentDirPath + "/"
	cd := configDirectory{}
	cd.currentDirPath = currentDirPath
	cd.parentDir = parentDir

	for _, file := range files {
		if !file.IsDir() {
			cd.filePaths = append(cd.filePaths, filePrefix+file.Name())
		}
		if file.IsDir() && file.Name() != ".git" { //todo add exclude group into config
			innercd := configDirectory{}
			innercd.currentDirPath = filePrefix + file.Name()
			innercd.parentDir = &cd
			cd.dirs = append(cd.dirs, innercd)
		}
	}
	//printConfigDirectory(cd)

	for _, dir := range cd.dirs {
		tree(dir.currentDirPath, &cd)
	}
	return cd
}

func printConfigDirectory(cd configDirectory) {
	fmt.Println(cd.currentDirPath)
	fmt.Println("\tdirectories:")
	for _, dir := range cd.dirs {
		fmt.Print("\t")
		fmt.Println(dir)
	}
	fmt.Println("\tfiles:")
	for _, file := range cd.filePaths {
		fmt.Print("\t")
		fmt.Println(file)
	}
	fmt.Print("\t")
	if cd.parentDir != nil {
		fmt.Println("parentDir:", cd.parentDir.currentDirPath)
		fmt.Print("\t")
	}
	fmt.Println("currentDir:", cd.currentDirPath)
	fmt.Println("----")
	fmt.Println()
}

func wipPrintDirWithTreeChars(cd *configDirectory) string {
	return ""
}
