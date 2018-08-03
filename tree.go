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

type config struct {
	path      string
	name      string
	parentDir *config
	subFiles  []*config
	isDir     bool
	prefix    string
}

func tree(file os.FileInfo, configs *[]config, parent config, prefix string) {
	c := config{}
	c.parentDir = &parent
	c.name = prefix + file.Name()
	if parent.path != "" {
		c.path = parent.path + "/" + file.Name()
	} else {
		c.path = file.Name()
	}
	c.prefix = prefix
	*configs = append(*configs, c)
	parent.subFiles = append(parent.subFiles, &c)


	if file.IsDir() && file.Name() != ".git" {
		c.isDir = true

		files, err := ioutil.ReadDir(c.path)
		if err != nil {
			log.Println(err)
		} else {
			for _, subfile := range files {
				tree(subfile, configs, c, "│   " + prefix)
			}
		}
	} else {
		c.isDir = false
	}
}

func buildTree(path string) {
	fileinfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	var configs []config
	c := config{}
	c.path = ""
	c.name = path
	c.prefix = ""
	tree(fileinfo, &configs, c, c.prefix)
	for _, conf := range configs {
		fmt.Println(conf.name)
	}
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
