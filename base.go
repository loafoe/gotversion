package gotversion

import (
	"fmt"

	version "github.com/hashicorp/go-version"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// Base is a (potential) base version
type Base struct {
	*version.Version
	Head     *object.Commit
	Branch   string
	Strategy Strategy
	Offset   int
	Tag      Tag
}

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

// BuildMetadata returns the offset
func (b Base) BuildMetadata() int {
	return b.Offset
}

// MajorMinorPatch returns the primary semver string
func (b Base) MajorMinorPatch() string {
	return fmt.Sprintf("%d.%d.%d", b.Major(), b.Minor(), b.Patch())
}

// FullBuildMetaData returns branch and SHA details
func (b Base) FullBuildMetaData() string {
	return fmt.Sprintf("Branch.%s.Sha.%s", b.Branch, b.Head.Hash.String())
}

// PreReleaseNumber returns the number of commits since the Base
func (b Base) PreReleaseNumber() int {
	return b.Offset
}

func (b Base) String() string {
	return fmt.Sprintf("semver=v%s branch=%s offset=%d", b.Version.String(), b.Branch, b.Offset)
}

// Major returns the Major semver element
func (b Base) Major() int {
	if b.Version == nil {
		return 0
	}
	return b.Version.Segments()[0]
}

// Minor returns the Minor semver element
func (b Base) Minor() int {
	if b.Version == nil {
		return 0
	}
	return b.Version.Segments()[1]
}

// Patch returns the Major semver element
func (b Base) Patch() int {
	if b.Version == nil {
		return 0
	}
	return b.Version.Segments()[2]
}

// Commit returns the hash of the head
func (b Base) Commit() string {
	return b.Head.Hash.String()
}

// CommitDate returns the commit date of head
func (b Base) CommitDate() string {
	return b.Head.Author.When.UTC().Format("2006-01-02")
}

// BranchName returns the branch name
func (b Base) BranchName() string {
	return b.Branch
}
