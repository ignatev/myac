package main

import (
	"fmt"
	"io/ioutil"
	"log"
//	"os"
//	"path/filepath"
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

	for _, file := range files {
		cd := configDirectory{}
		if file.IsDir() {
			cd.dirPaths = append(cd.dirPaths, file.Name())
		} else {
			cd.filePaths = append(cd.filePaths, file.Name())
		}
		cd.currentDirPath = currentDir
		cd.parentDirPath = parentDir

		printConfigDirectory(cd)
	}
	return "", nil

	/*
		var cds []configDirectory
		err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			var filePaths []string
			var dirPaths []string
			if f.IsDir() {
				dirPaths = append(dirPaths, f.Name())
				fmt.Println(dirPaths)
			}
			if !f.IsDir() {
				filePaths = append(filePaths, f.Name())
				fmt.Println(filePaths)
			}
			cd := configDirectory{dir, "", filePaths, dirPaths}
			fmt.Println(cd)
			cds = append(cds, cd)

			return nil
		})
		fmt.Println(cds)
		return "", err
	*/

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
