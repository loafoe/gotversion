package gotversion

import (
	"fmt"

	version "github.com/hashicorp/go-version"
)

// Strategy determines which part of semver to bump
type Strategy int

// Const definitions
const (
	Major Strategy = 0
	Minor Strategy = 1
	Patch Strategy = 2
)

func (strategy Strategy) String() string {
	names := [...]string{
		"Major",
		"Minor",
		"Patch"}
	if strategy < Major || strategy > Patch {
		return "Unknown"
	}
	return names[strategy]
}

// Bump the Tag version based on the strategy
func Bump(currentVersion *version.Version, strategy Strategy) (*version.Version, error) {
	segments := currentVersion.Segments()
	major, minor, patch := segments[0], segments[1], segments[2]
	switch strategy {
	case Major:
		major++
		minor = 0
		patch = 0
	case Minor:
		minor++
		patch = 0
	case Patch:
		patch++
	}
	newVersion, err := version.NewVersion(fmt.Sprintf("v%d.%d.%d", major, minor, patch))
	if err != nil {
		return nil, err
	}
	return newVersion, nil
}
