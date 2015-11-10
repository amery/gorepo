package manifest

type Projecter interface {
	Path() (string, error)
	Name() (string, error)
	Revision() (string, error)
	Remote() (string, error)
}
