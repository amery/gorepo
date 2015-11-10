package tree

import (
	"gopkg.in/libgit2/git2go.v22"
)

type Repo struct {
	Path       string
	Repository *git.Repository
}
