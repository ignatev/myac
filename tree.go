package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//└ ─ │ ├
const (
	newLine      = "\n"
	emptySpace   = "    "
	middleItem   = "├── "
	continueItem = "│   "
	lastItem     = "└── "
)

func tree(dir string) string {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Can't to access the path %q: %v\n", dir, err)
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walkgin the path %q: %v\n", dir, err)
	}
	return ""
}
