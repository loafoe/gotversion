package main

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
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

	headRef, err := r.Head()
	CheckIfError(err)

	fmt.Printf("Branch=%s\n", BranchName(r))
	fmt.Printf("Tag=%s\n", TagName(r))

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: headRef.Hash()})
	CheckIfError(err)

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		return nil
	})

	fmt.Println(headRef)
}
