package domain

import "sort"

type Spaces struct {
	store map[string]Space
}

func NewSpaces() *Spaces {
	return &Spaces{
		store: make(map[string]Space),
	}
}

func (s Spaces) Add(spacesToAdd ...Space) {
	for _, space := range spacesToAdd {
		s.store[space.GUID] = space
	}
}

func (s Spaces) Get(guid string) (Space, bool) {
	space, ok := s.store[guid]
	return space, ok
}

func (s Spaces) Delete(guid string) {
	delete(s.store, guid)
}

func (s *Spaces) Clear() {
	s.store = make(map[string]Space)
}

func (spaces Spaces) Len() int {
	return len(spaces.store)
}

func (spaces Spaces) Items() []interface{} {
	guids := sort.StringSlice([]string{})
	for _, space := range spaces.store {
		guids = append(guids, space.GUID)
	}

	sort.Sort(guids)

	var items []interface{}
	for _, guid := range guids {
		items = append(items, spaces.store[guid])
	}

	return items
}
