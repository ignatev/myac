package main

import (
	"fmt"
//	"os"
//	"path/filepath"
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
	partentDirPath []string
	filePaths      []string
	dirPaths       []string
}

func tree(dir string) (string, error) {
//	dirStructure := configDirectory{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			_, err := tree(file.Name())
			if err != nil {
				fmt.Println("error")
			}
		}
	}
	return "", nil
}
