package manifest

import (
	"gopkg.in/libgit2/git2go.v22"
	"log"
	"os"
)

type Manifest struct {
	logger *log.Logger
	repos  map[string]*git.Repository
}

func NewManifest(logger *log.Logger) *Manifest {
	m := Manifest{}
	m.repos = make(map[string]*git.Repository)
	if logger == nil {
		logger = log.New(os.Stderr, "", log.LstdFlags)
	}
	m.logger = logger

	return &m
}

func (m *Manifest) AddProjectByPath(path string, r *git.Repository) error {

	m.logger.Printf("Adding %s\n", path)
	m.repos[path] = r

	return nil
}
