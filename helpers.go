package gotversion

import (
	version "github.com/hashicorp/go-version"
	git "gopkg.in/src-d/go-git.v4"
)

// HeadTag returns the tag name of HEAD, if any
func HeadTag(r *git.Repository) (*Tag, error) {
	headRef, err := r.Head()
	if err != nil {
		return nil, err
	}
	if name := headRef.Name(); name.IsTag() {
		v, err := version.NewVersion(headRef.Name().Short())
		if err != nil {
			return nil, err
		}
		return &Tag{
			Version:   v,
			Reference: *headRef,
			Hash:      headRef.Hash().String(),
		}, nil
	}
	return nil, ErrNotATag
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
