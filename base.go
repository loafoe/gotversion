package gotversion

import (
	"fmt"

	version "github.com/hashicorp/go-version"
)

// Base is a (potential) base version
type Base struct {
	*version.Version
	Head     string
	Branch   string
	Strategy Strategy
	Offset   int
	Tag      Tag
}

// BaseCollection implements the Sortable interface
type BaseCollection []*Base

// Bump version
func (b *Base) Bump() error {
	strategy := Patch
	switch b.Branch {
	case "develop":
		strategy = Minor
	}
	newVersion, err := Bump(b.Version, strategy)
	if err != nil {
		return err
	}
	b.Version = newVersion
	return nil
}

// Semver where it all boils down to
func (b Base) Semver() string {
	label := b.PreReleaseLabel()
	if label == "" {
		return b.Version.String()
	}
	return fmt.Sprintf("%s-%s.%d", b.Version.String(), label, b.Offset)
}

// PreReleaseLabel is determined by the branch name
func (b Base) PreReleaseLabel() string {
	switch b.Branch {
	case "develop":
		return "alpha"
	case "master":
		return ""
	}
	return ""
}

// FullBuildMetaData returns branch and SHA details
func (b Base) FullBuildMetaData() string {
	return "Branch." + b.Branch + "Sha." + b.Head
}

// PreReleaseNumber returns the number of commits since the Base
func (b Base) PreReleaseNumber() int {
	return b.Offset
}

func (b Base) String() string {
	return fmt.Sprintf("semver=v%s branch=%s offset=%d", b.Version.String(), b.Branch, b.Offset)
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
