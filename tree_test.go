package main

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	assertCorrectDirStructure := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			fmt.Println(want)
			t.Errorf("\ngot:\n%s \nwant:\n%s", got, want)
		}
	}

	t.Run("one dir with one file", func(t *testing.T) {
		//cd := configDirectory{}
		//tree := tree(".filesystem-repo/service-1", cd)
		got := "service-1\n└── generic-service.yml"
		want := "service-1\n└── generic-service.yml"
		assertCorrectDirStructure(t, got, want)
	})

	t.Run("root dir should has empty parentDir", func(t *testing.T) {
		tree := tree(".filesystem-repo/service-1", nil)
		if tree.parentDir != nil {
			t.Errorf("\ngot:\n%s \nwant:\nnil", tree.parentDir)
		}
	})

	t.Run("non-root dir should has parentDir", func(t *testing.T) {
		cd := configDirectory{}
		cd.currentDirPath = "I'am parent"
		tree := tree(".filesystem-repo/service-1", &cd)
		got := tree.parentDir.currentDirPath
		want := "I'am parent"
		assertCorrectDirStructure(t, got, want)
	})


}
