package main

import (
	"github.com/amery/gorepo/manifest"
	"github.com/amery/gorepo/tree"
)

const (
	REPO_SCAN_DEPTH = 6
)

func main() {
	var depth uint = REPO_SCAN_DEPTH
	excludes := []string{".repo", "downloads"}
	m := manifest.NewManifest(nil)

	for r := range tree.Scan(".", depth, excludes) {
		m.AddProjectByPath(r.Path, r.Repository)
	}
}
