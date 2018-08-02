package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

type config struct {
	path      string
	name      string
	parentDir *config
	subFiles  []*config
	isDir     bool
	padding   int
}

func tree2(file os.FileInfo, configs *[]config, parent config, padding int) {

	c := config{}
	c.parentDir = &parent
	c.name = file.Name()
	if parent.path != "" {
		c.path = parent.path + "/" + file.Name()
	} else {
		c.path = file.Name()
	}
	c.padding = padding
	*configs = append(*configs, c)
	parent.subFiles = append(parent.subFiles, &c)


	if file.IsDir() && file.Name() != ".git" {
		c.isDir = true

		files, err := ioutil.ReadDir(c.path)
		if err != nil {
			log.Fatal(err)
		} else {
			for _, subfile := range files {
				tree2(subfile, configs, c, padding+1)
			}
		}
	} else {
		c.isDir = false
	}
}

func invokeTree2(path string) {
	fileinfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	var configs []config
	c := config{}
	c.path = ""
	c.name = path
	tree2(fileinfo, &configs, c, 0)
	for _, conf := range configs {
		pad := ""
		for p := 0; p < conf.padding*2; p++ {
			pad = " " + pad
		}
		fmt.Println(pad + conf.name)
	}
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

func treeRefactor(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
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
