package keys

type Builder interface {
	Build() Entity

	SetSize(size SizeInBits) Builder
}

type Generator interface {
	NewBuilder() Builder
}
