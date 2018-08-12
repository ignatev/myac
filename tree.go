package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	newline      = "\n"
	emptyspace   = "    "
	middleitem   = "├── "
	continueitem = "│   "
	lastitem     = "└── "
)

type tree struct {
	path, name string
	children   []*tree
	parent     *tree
	fileinfo   os.FileInfo
	//todo add fileinfo for paths and isDir() func
}

func buildtree(path string, parent *tree) *tree { //todo use fileinfo
	var current tree
	var children []*tree
	current.name = path
	current.parent = parent
	current.path = parent.path + "/" + current.name
	fileinfo, err := os.Stat(current.path)
	if err != nil {
		log.Println(err)
	}
	if fileinfo.IsDir() && fileinfo.Name() != ".git" {
		dir, err := ioutil.ReadDir(current.path)
		if err != nil {
			log.Println(err)
		}
		for _, file := range dir {
			child := buildtree(file.Name(), &current)
			children = append(children, child)
		}
		current.children = children
	}
	return &current
}

func runbuildtree(path string) {
	var rootDir tree
	rootDir.path = filepath.Dir(path)

	tree := buildtree(filepath.Base(path), &rootDir)
	for _, d := range rendertree(tree) {
		fmt.Println(d)
	}
}

func rendertree(tree *tree) []string {
	var result []string
	result = append(result, tree.name)

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
