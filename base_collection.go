package gotversion

// BaseCollection implements the Sortable interface
type BaseCollection []*Base

func (v BaseCollection) Len() int {
	return len(v)
}

func (v BaseCollection) Less(i, j int) bool {
	return v[i].Version.LessThan(v[j].Version)
}

func (v BaseCollection) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
