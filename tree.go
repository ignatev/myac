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

type dirtree struct {
	path, name	string
	children	[]*dirtree
	parent		*dirtree
	isDir		bool
}

func fillTree(root string, parent *dirtree, isDir bool) dirtree {
	var current dirtree
	var children []*dirtree
	current.path = parent.path + "/" + root
	current.name = root
	current.parent = parent
	current.isDir = isDir

	if isDir {
		dir, err := ioutil.ReadDir(root)
		if err != nil {
			// log
		}
		for _, file := range dir {
			if file.IsDir() {
				child := fillTree(file.Name(), &current, true)
				children = append(children, &child)
			}
		}
	}
	current.children = children
	return current
}

func runFillTree(root string) {
	var rootDir dirtree
	rootDir.path = ""
	rootDir.name = root
	dir, err := os.Stat(root)
	if err != nil {
		// log
	}
	if dir.IsDir() {
		result := fillTree(root, &rootDir, true)
		fmt.Println(result)
	}
}