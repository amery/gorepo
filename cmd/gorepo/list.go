package main

import (
	"fmt"
	"github.com/amery/gorepo/tree"
)

const (
	REPO_SCAN_DEPTH = 6
)

func repo_list() int {
	var depth uint = REPO_SCAN_DEPTH
	excludes := []string{".repo"}
	ret := 1

	for r := range tree.Scan(".", depth, excludes) {
		fmt.Printf("%s\n", r.Path)
		ret = 0
	}

	return ret
}

func init() {
	repo_commands["list"] = repo_list
}
