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
	url, path, name string
	children        []*tree
	parent          *tree
	mapping         *map[string][]string
}

func buildtree(path string, parent *tree) *tree {
	var current tree
	var children []*tree
	current.name = path
	current.parent = parent
	current.path = parent.path + "/" + current.name
	current.mapping = parent.mapping

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
	if !fileinfo.IsDir() {
		current.url = parent.name
		c := current.mapping

		v := (*c)[parent.name]
		v = append(v, current.path)
		(*c)[parent.name] = v

	}
	return &current
}

func treebuilder(path string) {
	var rootDir tree
	m := make(map[string][]string)
	rootDir.path = filepath.Dir(path)
	rootDir.mapping = &m
	tree := buildtree(filepath.Base(path), &rootDir)

	for _, d := range rendertree(tree) {
		fmt.Println(d)
	}
}

func rendertree(tree *tree) []string {
	var result []string
	var mapping string
	if tree.url != "" {
		mapping = " >>> " + "http://localhost:8888/" + tree.url
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
