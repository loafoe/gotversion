package gotversion

import (
	"fmt"

	version "github.com/hashicorp/go-version"
)

// Base is a (potential) base version
type Base struct {
	*version.Version
	Branch   string
	Strategy string
	Offset   int64
}

// BaseCollection implements the Sortable interface
type BaseCollection []*Base

func (b Base) String() string {
	return fmt.Sprintf("semver=%s branch=%s offset=%d", b.Version.String(), b.Branch, b.Offset)
}

func (v BaseCollection) Len() int {
	return len(v)
}

func (v BaseCollection) Less(i, j int) bool {
	return v[i].Version.LessThan(v[j].Version)
}

func (v BaseCollection) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
