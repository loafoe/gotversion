package gotversion

import (
	version "github.com/hashicorp/go-version"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// SemverTags returns only valid semver tags
func SemverTags(r *git.Repository) (*[]Tag, error) {
	tagrefs, err := r.Tags()
	if err != nil {
		return nil, err
	}
	list := []Tag{}
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		tagName := t.Name().Short()
		if v, err := version.NewVersion(tagName); err == nil {
			tag := Tag{
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
