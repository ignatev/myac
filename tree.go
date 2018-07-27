package main

import (
	"path/filepath"
	"os"
)

//└ ─ │ ├
const (
	newLine			= "\n"
	emptySpace  	= "    "
	middleItem		= "├── "
	continueItem	= "│   "
	lastItem 		= "└── "
)

func tree(dir string) string {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			err.Errorf("Fail to access a path %q: %v\n", dir, err)
		}
	})

}
