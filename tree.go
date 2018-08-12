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
		current.path = path
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
	tree := fillTree(root, &rootDir)
	//	fmt.Println("run", tree, root)
	for _, d := range renderTree(tree) {
		fmt.Println(d)
	}
}

func printtree(tree *dirtree) {
	fmt.Println(tree.name)
	fmt.Println(tree.path)
	if len(tree.children) != 0 {
		for _, tree := range tree.children {
			printtree(tree)
		}
	}
}

func renderTree(tree *dirtree) []string {
	var result []string
	result = append(result, tree.name)

	fmt.Println("000", result)
		for i, child := range tree.children {
			fmt.Println(result)
			subtr := renderTree(child)
			fmt.Println(subtr)
			fmt.Println(result)
			if i == len(tree.children)-1 {
				result = lastsubtree(result, subtr)
			} else {
				result = subtree(result, subtr)
			}
		}
	fmt.Println("999", result)
	return result
}

func subtree(result, subtr []string) []string {
	fmt.Println("1", result, subtr)
	result = append(result, middleItem+subtr[0])
	fmt.Println("2", result, subtr)

	for _, child := range subtr[1:] {
		fmt.Println("3", result)
		fmt.Println("4", child)
		result = append(result, continueItem+child)
		fmt.Println("5", result)
	}

	fmt.Println("6", result)
	return result
}

func lastsubtree(result, subtr []string) []string {
	fmt.Println("1-1", result, subtr)
	result = append(result, lastItem+subtr[0])
	fmt.Println("1-2", result, subtr)
	for _, child := range subtr[1:] {
		result = append(result, emptySpace+child)
	}
	return result
}
