package keys

type Loader interface {
	Load(p *PEM) (Entity, error)
}
