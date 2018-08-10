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
				tree(subfile, configs, c, "│   "+prefix)
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
	path, name string
	children   []*dirtree
	parent     *dirtree
}

func fillTree(path string, parent *dirtree) *dirtree {
	var current dirtree
	var children []*dirtree
	current.name = path
	current.parent = parent
	if parent.path == "" {
		current.path = parent.name + "/" + path
	} else {
		current.path = parent.path + "/" + path
	}
//	fmt.Println("current path:", current.path)

	fileinfo, err := os.Stat(current.path)
	if err != nil {
		fmt.Println(err)
	}
	if fileinfo.IsDir() && fileinfo.Name() != ".git" {

	dir, err := ioutil.ReadDir(current.path)
	if err != nil {
//		fmt.Println(err)
	}
	for _, file := range dir {
		child := fillTree(file.Name(), &current)
		children = append(children, child)
	}
	current.children = children
	}
	return &current
}

func runFillTree(root string) {
	var rootDir dirtree
	rootDir.path = ""
	rootDir.name = root
//	var children []*dirtree
//	rootDir.children = children

dir, err := ioutil.ReadDir(root)
	fmt.Println(dir)
	if err != nil {

	}
	for _, d := range dir {
		fillTree(d.Name(), &rootDir)
	}
	printtree(&rootDir)

}

func printtree(tree *dirtree) {
	fmt.Println(tree.name)
	fmt.Println(tree.children)
	if len(tree.children) != 0 {
		for _, tree := range tree.children {
			printtree(tree)
		}
	}
}

func renderTree(tree *dirtree) []string {
	var result []string
	result = append(result, tree.name)
	for i, child := range tree.children {
//		result = append(result, renderTree(child)...)
		if i == len(tree.children)-1 {
			result = append(result, lastsubtree(child)...)
		} else {
			result = append(result, subtree(child)...)
		}
	}
	for _, p := range result {
		fmt.Println(p)
	}
	return result
}

func subtree(subtree *dirtree) []string {
	var result []string
	result = append(result, middleItem+subtree.name)
	for _, child := range subtree.children {
		result = append(result, continueItem+child.name)
	}
	return result
}

func lastsubtree(subtree *dirtree) []string {
	var result []string
	result = append(result, lastItem+subtree.name)
	for _, child := range subtree.children {
		result = append(result, emptySpace+child.name)
	}
	return result
}
