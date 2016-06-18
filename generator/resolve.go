package generator

import (
	"fmt"
)

// A label represents a label of a build target in Bazel.
type label struct {
	repo, pkg, name string
}

func (l label) String() string {
	if l.repo != "" {
		return fmt.Sprintf("@%s//%s:%s", l.repo, l.pkg, l.name)
	}
	return fmt.Sprintf("//%s:%s", l.pkg, l.name)
}

// A labelResolver resolves a Go importpath into a label in Bazel.
type labelResolver interface {
	resolve(importpath string) (label, error)
}

type resolverFunc func(importpath string) (label, error)

func (f resolverFunc) resolve(importpath string) (label, error) {
	return f(importpath)
}
