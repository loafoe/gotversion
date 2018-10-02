package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
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
}
