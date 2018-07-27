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
		got := tree(".filesystem-repo/service-1")
		want := "service-1\n└── generic-service.yml"
		assertCorrectDirStructure(t, got, want)
	})
}
