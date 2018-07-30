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
	partentDirPath string
	filePaths      []string
	dirPaths       []string
}

func tree(dir string) (string, error) {
	
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		fmt.Println(file.IsDir())
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
