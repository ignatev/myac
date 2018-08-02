package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
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
	result := cd.currentDirPath + "\n"
	dirs := cd.dirs
	files := cd.filePaths
	for i, dir := range dirs {
		result = dirContent(result, dir.currentDirPath, i, len(dirs))
	}
	for i, file := range files {
		result = dirContent(result, file, i, len(files))
	}
	return result
}

func dirContent(result, path string, i, l int) string {
	p := fileName(path)
	if i < l-1 {
		result += "├── " + p + "\n"
	} else {
		result += "└── " + p + "\n"
	}
	return result
}

func fileName(fileName string) string {
	segs := strings.Split(fileName, "/")
	return segs[len(segs)-1]
}
