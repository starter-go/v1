package keys

type Entity interface {
	Algorithm() Algorithm

	Size() SizeInBits

	GetDriver() Driver
}

type KeyEntity interface {
	Entity
}
