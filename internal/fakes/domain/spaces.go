package domain

type Spaces struct {
	store map[string]Space
}

func NewSpaces() *Spaces {
	return &Spaces{
		store: make(map[string]Space),
	}
}

func (s Spaces) Get(guid string) (Space, bool) {
	space, ok := s.store[guid]
	return space, ok
}

func (s Spaces) Add(space Space) {
	s.store[space.GUID] = space
}

func (s Spaces) Delete(guid string) {
	delete(s.store, guid)
}

func (s *Spaces) Clear() {
	s.store = make(map[string]Space)
}
