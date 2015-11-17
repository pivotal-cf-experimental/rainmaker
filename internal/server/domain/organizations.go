package domain

import "sort"

type Organizations struct {
	store map[string]Organization
}

func NewOrganizations() *Organizations {
	return &Organizations{
		store: make(map[string]Organization),
	}
}

func (o Organizations) Add(orgsToAdd ...Organization) {
	for _, org := range orgsToAdd {
		o.store[org.GUID] = org
	}
}

func (o Organizations) Get(guid string) (Organization, bool) {
	org, ok := o.store[guid]
	return org, ok
}

func (o Organizations) Delete(guid string) {
	delete(o.store, guid)
}

func (o *Organizations) Clear() {
	o.store = make(map[string]Organization)
}

func (o Organizations) Len() int {
	return len(o.store)
}

func (o Organizations) Items() []interface{} {
	guids := sort.StringSlice([]string{})
	for _, org := range o.store {
		guids = append(guids, org.GUID)
	}

	sort.Sort(guids)

	var items []interface{}
	for _, guid := range guids {
		items = append(items, o.store[guid])
	}

	return items
}
