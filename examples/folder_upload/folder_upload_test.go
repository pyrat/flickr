package main

import (
	"testing"
)

func TestGetFilePaths(t *testing.T) {
	paths := getFilePaths("./")
	if len(paths) == 0 {
		t.Error("Unable to get the filepaths")
	}
}
