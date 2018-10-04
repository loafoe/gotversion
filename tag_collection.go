package gotversion

// TagCollection implements the Sortable interface
type TagCollection []*Tag

func (v TagCollection) Len() int {
	return len(v)
}

func (v TagCollection) Less(i, j int) bool {
	return v[i].Version.LessThan(v[j].Version)
}

func (v TagCollection) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
