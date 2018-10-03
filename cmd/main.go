package main

import (
	"fmt"
	"sort"

	version "github.com/hashicorp/go-version"
	"github.com/loafoe/gotversion"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	"os"
)

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// BranchName returns the branch name which we're on
func BranchName(r *git.Repository) string {
	headRef, err := r.Head()
	if err != nil {
		return ""
	}
	if name := headRef.Name(); name.IsBranch() {
		return name.Short()
	}
	return ""
}

// TagName returns the tag name of HEAD, if any
func TagName(r *git.Repository) string {
	headRef, err := r.Head()
	if err != nil {
		return ""
	}
	if name := headRef.Name(); name.IsTag() {
		return "tag:" + name.Short()
	}
	return "notag"
}

// SemverTags returns only valid semver tags
func SemverTags(r *git.Repository) (*[]gotversion.Tag, error) {
	tagrefs, err := r.Tags()
	if err != nil {
		return nil, err
	}
	list := []gotversion.Tag{}
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		tagName := t.Name().Short()
		if v, err := version.NewVersion(tagName); err == nil {
			tag := gotversion.Tag{
				Version:   v,
				Reference: *t,
				Hash:      t.Hash().String(),
			}
			obj, err := r.TagObject(t.Hash())
			if err == nil {
				tag.IsAnnoted = true
				tag.Hash = obj.Target.String()
			}
			list = append(list, tag)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: getver <path>\n")
		os.Exit(1)
	}

	path := os.Args[1]

	r, err := git.PlainOpen(path)
	CheckIfError(err)

	tagrefs, err := r.Tags()
	CheckIfError(err)

	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		fmt.Println(t)
		return nil
	})
	CheckIfError(err)

	semverTags, err := SemverTags(r)
	CheckIfError(err)

	fmt.Printf("%d SemverTags found\n", len(*semverTags))
	if len(*semverTags) > 0 {
		for _, t := range *semverTags {
			fmt.Println(t)
		}
	}

	headRef, err := r.Head()
	CheckIfError(err)

	branchName := BranchName(r)
	fmt.Printf("Branch=%s\n", branchName)
	fmt.Printf("Tag=%s\n", TagName(r))

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: headRef.Hash()})
	CheckIfError(err)

	// ... just iterates over the commits, printing it
	baseVersions := gotversion.BaseCollection{}
	var offset int64
	err = cIter.ForEach(func(c *object.Commit) error {
		for _, t := range *semverTags {
			if t.Hash == c.Hash.String() {
				fmt.Println("Adding baseVersion...")
				baseVersions = append(baseVersions, &gotversion.Base{
					Branch:   branchName,
					Strategy: "w0t",
					Version:  t.Version,
					Offset:   offset,
				})
			}
		}
		offset++
		return nil
	})
	sort.Sort(baseVersions)

	fmt.Println(baseVersions[len(baseVersions)-1])
}