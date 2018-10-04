package gotversion

import (
	"fmt"
	"regexp"
	"strings"

	version "github.com/hashicorp/go-version"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var featureRegexp *regexp.Regexp

func init() {
	featureRegexp = regexp.MustCompile(`^feature/`)
}

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
	if b.Tag.Hash == "" { // No tags found!
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
	fmtString := "%s-%s.%d"
	return fmt.Sprintf(fmtString, b.Version.String(), label, b.Offset)
}

// HeadTag returns true if the tag is on the branch head
// Typically this means we should not Bump() the version
func (b Base) HeadTag() bool {
	return b.Offset == 0
}

// FullSemver returns Semver with offset in case no tags are there
func (b Base) FullSemver() string {
	if b.IsMasterBranch() && !b.HeadTag() {
		return fmt.Sprintf("%s+%d", b.Version.String(), b.Offset)
	}
	if b.HeadTag() {
		return b.Semver()
	}
	if label := b.PreReleaseLabel(); label != "" {
		return fmt.Sprintf("%s-%s+%d", b.Version.String(), label, b.Offset)
	}
	// Tag is on the head
	return b.Semver()
}

// IsMasterBranch returns true if we are on the master branch
func (b Base) IsMasterBranch() bool {
	return b.Branch == "master"
}

// IsDevelopBranch return true if we are on the develop branch
func (b Base) IsDevelopBranch() bool {
	return b.Branch == "develop"
}

// IsFeatureBranch returns true if we are on a feature branch
func (b Base) IsFeatureBranch() bool {
	return featureRegexp.FindStringSubmatch(b.Branch) != nil
}

// PreReleaseLabel is determined by the branch name
func (b Base) PreReleaseLabel() string {
	if b.IsDevelopBranch() {
		return "alpha"
	}
	if b.IsMasterBranch() {
		return ""
	}
	if b.IsFeatureBranch() {
		return strings.Replace(b.Branch, "feature/", "", 1) + ".1"
	}
	return b.Branch
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
func (b Base) PreReleaseNumber() string {
	return ""
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
