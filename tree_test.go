package main

import "testing"

func TestTree(t *testing.T) {
	got := tree("my_path")
	want := "my_path"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}