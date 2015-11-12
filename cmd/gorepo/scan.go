package main

import (
	"github.com/amery/gorepo/manifest"
	"github.com/amery/gorepo/tree"
)

const (
	REPO_SCAN_DEPTH = 6
)

func repo_scan() int {
	var depth uint = REPO_SCAN_DEPTH
	excludes := []string{".repo", "downloads"}
	m := manifest.NewManifest(nil)

	for r := range tree.Scan(".", depth, int(depth/2), excludes) {
		m.AddProjectByPath(r.Path, r.Repository)
	}

	return 0
}

func init() {
	repo_commands["scan"] = repo_scan
}
