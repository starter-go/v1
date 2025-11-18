package keys

type Driver interface {
	GetRegistration() *DriverRegistration
}

type KeyDriver interface {
	Driver

	GetLoader() Loader

	GetSaver() Saver

	GetGenerator() Generator

	GetCipher() Cipher
}

type DriverRegistration struct {
	Name string

	Algorithm Algorithm

	Enabled bool

	Priority int

	Driver Driver
}

type DriverRegistry interface {
	RegisterDriver(dr Driver) error
}

type DriverManager interface {
	FindDriver(alg Algorithm) (Driver, error)
}
