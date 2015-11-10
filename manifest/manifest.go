package manifest

import (
	"fmt"
)

type Manifest struct{}

func (m Manifest) AddProjecter(p Projecter) error {
	path, err := p.Path()
	if err != nil {
		fmt.Printf("! %s %v\n", err, p)
		return err
	}

	fmt.Printf("- %s\n", path)
	return nil
}
