package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	newline      = "\n"
	emptyspace   = "    "
	middleitem   = "├── "
	continueitem = "│   "
	lastitem     = "└── "
)

type tree struct {
	url, name			string
	relpath, abspath	string
	children        	[]*tree
	parent          	*tree
	mapping         	*map[string][]string
	isfile				bool
}

func buildtree(path string, parent *tree) *tree {
	var current tree
	var children []*tree
	current.name = path
	current.parent = parent
	current.relpath = parent.relpath + "/" + current.name
	current.abspath = parent.abspath + "/" + current.name
	current.mapping = parent.mapping

	fileinfo, err := os.Stat(current.abspath)
	if err != nil {
		log.Println(err)
	}
	if fileinfo.IsDir() && fileinfo.Name() != ".git" {
		dir, err := ioutil.ReadDir(current.abspath)
		if err != nil {
			log.Println(err)
		}
		for _, file := range dir {
			child := buildtree(file.Name(), &current)
			children = append(children, child)
		}
		current.children = children
	}
	if !fileinfo.IsDir() {
		current.url = parent.name
		current.isfile = true
		c := current.mapping

		v := (*c)[parent.name]
		v = append(v, current.abspath)
		(*c)[parent.name] = v
	}

	return &current
}

func treebuilder(path string) {
	var rootDir tree
	m := make(map[string][]string)
	rootDir.relpath = ""
	rootDir.abspath = filepath.Dir(path)
	rootDir.mapping = &m

	tree := buildtree(filepath.Base(path), &rootDir)
	for _, d := range rendertree(tree) {
		fmt.Println(d)
	}
}

func rendertree(tree *tree) []string {
	var result []string
	var mapping string
	if tree.isfile {
		m := (*tree.mapping)[tree.parent.name]
		prefix := tree.parent.relpath
		if len(m) == 1 {
			tree.url = prefix
		} else {
			tree.url = prefix + "/" + strings.TrimSuffix(tree.name, filepath.Ext(tree.name))
		}
		mapping = " >>> " + "http://localhost:8888" + tree.url
	}

	result = append(result, tree.name+mapping)

	for i, child := range tree.children {
		subtr := rendertree(child)
		if i == len(tree.children)-1 {
			result = lastsubtree(result, subtr)
		} else {
			result = subtree(result, subtr)
		}
	}

	return result
}

func subtree(result, subtr []string) []string {
	result = append(result, middleitem+subtr[0])
	for _, child := range subtr[1:] {
		result = append(result, continueitem+child)
	}

	return result
}

func lastsubtree(result, subtr []string) []string {
	result = append(result, lastitem+subtr[0])
	for _, child := range subtr[1:] {
		result = append(result, emptyspace+child)
	}

	return result
}
