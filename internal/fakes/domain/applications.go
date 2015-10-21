package domain

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
