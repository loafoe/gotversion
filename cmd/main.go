package main

import (
	"flag"
	"fmt"
	"sort"

	version "github.com/hashicorp/go-version"
	"github.com/loafoe/gotversion"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"

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

func main() {

	showJSON := flag.Bool("json", false, "output JSON")
	showVSO := flag.Bool("vso", true, "output VSO")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Printf("usage: gotversion [-json|-vso] path\n")
		os.Exit(1)
	}
	path := args[0]

	r, err := git.PlainOpen(path)
	CheckIfError(err)

	semverTags, err := gotversion.SemverTags(r)
	CheckIfError(err)

	//fmt.Printf("%d SemverTags found\n", len(*semverTags))
	tagCollection := gotversion.TagCollection{}
	for _, t := range *semverTags {
		newTag := t
		tagCollection = append(tagCollection, &newTag)
	}
	sort.Sort(sort.Reverse(tagCollection))

	headRef, err := r.Head()
	CheckIfError(err)

	branchName := gotversion.BranchName(r)
	//fmt.Printf("Branch=%s\n", branchName)
	//fmt.Printf("Tag=%s\n", TagName(r))

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: headRef.Hash()})
	CheckIfError(err)

	headCommit, err := r.CommitObject(headRef.Hash())
	CheckIfError(err)

	// ... just iterates over the commits, printing it
	var baseVersion *gotversion.Base
	var offset int
	err = cIter.ForEach(func(c *object.Commit) error {
		for _, t := range tagCollection {
			if t.Hash == c.Hash.String() {
				baseVersion = &gotversion.Base{
					Head:     headCommit,
					Branch:   branchName,
					Strategy: gotversion.Patch,
					Version:  t.Version,
					Offset:   offset,
					Tag:      *t,
				}
				break
			}
		}
		if baseVersion != nil {
			// Stop immediately
			return storer.ErrStop
		}
		offset++
		return nil
	})

	if baseVersion == nil { // No semver tags found
		version, _ := version.NewVersion("v0.0.0")
		baseVersion = &gotversion.Base{
			Version:  version,
			Head:     headCommit,
			Branch:   branchName,
			Strategy: gotversion.Minor,
			Offset:   offset,
			Tag:      gotversion.Tag{},
		}
	}
	if !baseVersion.HeadTag() {
		baseVersion.Bump()
	}
	if *showJSON {
		*showVSO = false
		gotversion.OutputJSON(baseVersion)
	}
	if *showVSO {
		gotversion.OutputVSO(baseVersion)
	}
}
