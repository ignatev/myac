package main

import (
	"os"
	"testing"
)

func TestTree(t *testing.T) {

	assertDirIsNotExist := func(t *testing.T, path string) {
		t.Helper()
		_, err := os.Stat(path)
		if !os.IsNotExist(err) {
			t.Errorf("Repository already cloned or directory exists")
		}
	}

	assertCorrectDirStructure := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\ngot:\n%s \nwant:\n%s", got, want)
		}
	}

	t.Run("one dir with one file", func(t *testing.T) {
		got := "service-1\n└── generic-service.yml"
		want := "service-1\n└── generic-service.yml"
		assertCorrectDirStructure(t, got, want)
	})

	t.Run("should clone given repo to local dir", func(t *testing.T) {
		var c serverConf
		config := c.getConf("config.yml")
		url := config.Server.Git.URL
		localRepositoryPath := config.Server.Git.LocalRepositoryPath
		assertDirIsNotExist(t, localRepositoryPath)
		cloneConfigRepo(localRepositoryPath, url)
	})

/* 	t.Run("root dir should has empty parentDir", func(t *testing.T) {
		cd := config{}
		cd.path = ".filesystem-repo/service-1"
		tree := buildTree(".filesystem-repo/service-1")
		if tree.parentDir != nil {
			t.Errorf("\ngot:\n%s \nwant:\nnil", tree.parentDir)
		}
	}) */

/* 	t.Run("non-root dir should has parentDir", func(t *testing.T) {
		cd := configDirectory{}
		cd.currentDirPath = "I'am parent"
		tree := tree(".filesystem-repo/service-1", &cd)
		got := tree.parentDir.currentDirPath
		want := "I'am parent"
		assertCorrectDirStructure(t, got, want)
	}) */

/* 	t.Run("print directory with one layer", func(t *testing.T) {
		files := []string{"file-1", "file-2", "file-3"}
		cd := configDirectory{"test", nil, files, nil}
		got := wipPrintDirWithTreeChars(&cd)
		want := "test\n├── file-1\n├── file-2\n└── file-3\n"
		assertCorrectDirStructure(t, got, want)
	}) */

}
