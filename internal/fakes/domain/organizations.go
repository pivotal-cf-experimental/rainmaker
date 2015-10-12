package domain

type Organizations struct {
	store map[string]Organization
}

func NewOrganizations() *Organizations {
	return &Organizations{
		store: make(map[string]Organization),
	}
}

func (o Organizations) Get(guid string) (Organization, bool) {
	org, ok := o.store[guid]
	return org, ok
}

func (o Organizations) Add(org Organization) {
	o.store[org.GUID] = org
}

func (o Organizations) Delete(guid string) {
	delete(o.store, guid)
}

func (o *Organizations) Clear() {
	o.store = make(map[string]Organization)
}
