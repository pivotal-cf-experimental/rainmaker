package domain

type Buildpacks struct {
	store map[string]Buildpack
}

func NewBuildpacks() *Buildpacks {
	return &Buildpacks{
		store: make(map[string]Buildpack),
	}
}

func (b Buildpacks) Add(bp Buildpack) {
	b.store[bp.GUID] = bp
}

func (b Buildpacks) Get(guid string) (Buildpack, bool) {
	bp, ok := b.store[guid]
	return bp, ok
}

func (b Buildpacks) Remove(guid string) {
	delete(b.store, guid)
}
