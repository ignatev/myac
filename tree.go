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
	dirPaths       []configDirectory
}

func tree(currentDir configDirectory, parentDir string) (string, error) {
	fmt.Println("start currentDir for outer dir:")
	fmt.Println(currentDir)
	fmt.Println("end currentDir for outer dir:")
	files, err := ioutil.ReadDir(currentDir.currentDirPath)
	if err != nil {
		log.Fatal(err)
	}

	filePrefix := currentDir.currentDirPath + "/"
	cd := configDirectory{}
	cd.currentDirPath = currentDir.currentDirPath

	for _, file := range files {
		if !file.IsDir() {
			cd.filePaths = append(cd.filePaths, filePrefix + file.Name())
		}
		if file.IsDir() && file.Name() != ".git" { //todo add exclude group into config
			innercd := configDirectory{}
			innercd.currentDirPath = filePrefix + file.Name()
			fmt.Println("start currentDir for inner dir:")
			fmt.Println(currentDir)
			fmt.Println("end currentDir for inner dir:")
			innercd.parentDir = &currentDir
			fmt.Println(innercd.parentDir.currentDirPath)
			cd.dirPaths = append(cd.dirPaths, innercd)
		}
	}
	printConfigDirectory(cd)

	for _, dir := range cd.dirPaths {
		tree(dir, cd.currentDirPath)
	}

	return "", nil

}

func printConfigDirectory(cd configDirectory) {
	fmt.Println(cd.currentDirPath)
	fmt.Println("\tdirectories:")
	for _, dir := range cd.dirPaths {
		fmt.Print("\t")
		fmt.Println(dir)
	}
	fmt.Println("\tfiles:")
	for _, file := range cd.filePaths {
		fmt.Print("\t")
		fmt.Println(file)
	}
	fmt.Print("\t")
	fmt.Println(cd)
	fmt.Println(cd.parentDir)
	if cd.parentDir != nil {
		
		fmt.Println("parentDir:", &cd.parentDir.currentDirPath)
		fmt.Print("\t")
	}
	fmt.Println("currentDir:", cd.currentDirPath)
	fmt.Println("----")
	fmt.Println()
}
