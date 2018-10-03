package gotversion

import (
	"fmt"

	version "github.com/hashicorp/go-version"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Tag holds our representation of a Tag (both annoted and lightweight)
type Tag struct {
	*version.Version
	Reference plumbing.Reference
	IsAnnoted bool
	Hash      string
}

func (t Tag) String() string {
	return fmt.Sprintf("v%s [annotated=%t] %s", t.Version, t.IsAnnoted, t.Hash)
}
