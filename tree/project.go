package tree

import (
	"errors"
)

func (r Repo) Path() (string, error) {
	return r.path, nil
}

func (r Repo) Name() (string, error) {
	return "", errors.New("Not Implemented")
}

func (r Repo) Remote() (string, error) {
	return "", errors.New("Not Implemented")
}

func (r Repo) Revision() (string, error) {
	return "", errors.New("Not Implemented")
}
