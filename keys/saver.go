package keys

type Saver interface {
	Save(key Entity) (*PEM, error)
}
