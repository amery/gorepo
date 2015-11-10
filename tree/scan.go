package tree

import (
	"gopkg.in/libgit2/git2go.v22"
	"os"
	"path"
	"sync"
)

const (
	SUPPORTS_NESTED = true
)

func load_repo(dirname string, ch chan<- Repo) error {
	r, err := git.OpenRepository(dirname)
	if err != nil {
		return err
	}

	ch <- Repo{dirname, r}
	return nil
}

func readdirnames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}

	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	return names, nil
}

func scan_children(dirname string, names []string, result chan<- Repo,
	depth uint, exclude map[string]bool) error {
	var wg sync.WaitGroup

	for _, fn := range names {
		name := path.Join(dirname, fn)

		if !exclude[name] {
			wg.Add(1)
			go func() {
				scan(name, result, depth, exclude)
				defer wg.Done()
			}()
		}
	}

	wg.Wait()
	return nil
}

func scan(name string, result chan<- Repo,
	depth uint, exclude map[string]bool) error {
	// don't follow symlinks
	info, err := os.Lstat(name)
	if err != nil {
		return err
	} else if info.Name() == ".git" {
		// regular working copy
		if SUPPORTS_NESTED {
			// skip to avoid duplication
			// TODO: be smarter, compare with parent
			return nil
		} else {
			return load_repo(name, result)
		}
	} else if depth > 0 && info.IsDir() {
		err := load_repo(name, result)
		if SUPPORTS_NESTED || err != nil {
			// not a bare repo, walk through it
			names, err := readdirnames(name)
			if err != nil {
				return err
			}

			return scan_children(name, names, result, depth-1, exclude)
		}
	}
	return nil
}

func Scan(dirname string, depth uint, exclude []string) chan Repo {
	result := make(chan Repo, 10)
	m := map[string]bool{}

	for _, s := range exclude {
		m[s] = true
	}

	go func() {
		scan(dirname, result, depth, m)
		close(result)
	}()
	return result
}
