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

func (b Buildpacks) Get(guid string) Buildpack {
	return b.store[guid]
}
