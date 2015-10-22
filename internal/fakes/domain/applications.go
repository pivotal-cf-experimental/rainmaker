package domain

import "sort"

type Applications struct {
	store map[string]Application
}

func NewApplications() *Applications {
	return &Applications{
		store: make(map[string]Application),
	}
}

func (a Applications) Add(appsToAdd ...Application) {
	for _, app := range appsToAdd {
		a.store[app.GUID] = app
	}
}

func (a Applications) Get(guid string) (Application, bool) {
	app, ok := a.store[guid]
	return app, ok
}

func (a Applications) Delete(guid string) {
	delete(a.store, guid)
}

func (a *Applications) Clear() {
	a.store = make(map[string]Application)
}

func (a Applications) Len() int {
	return len(a.store)
}

func (a Applications) Items() []interface{} {
	guids := sort.StringSlice([]string{})
	for _, app := range a.store {
		guids = append(guids, app.GUID)
	}

	sort.Sort(guids)

	var items []interface{}
	for _, guid := range guids {
		items = append(items, a.store[guid])
	}

	return items
}
