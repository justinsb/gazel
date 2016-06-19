package generator

import (
	"fmt"
	"strings"
)

// flatResolver resolves go_library labels within the same repository as
// the one of goPrefix, assuming all rules are defined in a single BUILD file.
type flatResolver struct {
	goPrefix string
}

func (r flatResolver) resolve(importpath, _ string) (label, error) {
	if importpath == r.goPrefix {
		return label{name: "go_default_library", relative: true}, nil
	}

	if prefix := r.goPrefix + "/"; strings.HasPrefix(importpath, prefix) {
		return label{
			name:     strings.TrimPrefix(importpath, prefix),
			relative: true,
		}, nil
	}

	return label{}, fmt.Errorf("importpath %q does not start with goPrefix %q", importpath, r.goPrefix)
}
