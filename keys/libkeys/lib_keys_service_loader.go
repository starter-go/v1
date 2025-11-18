package libkeys

import "github.com/starter-go/v1/keys"

type Loader struct{}

// Load implements keys.ServiceLoader.
func (inst *Loader) Load() keys.Service {

	ser := new(innerLibKeysService)

	return ser
}

func (inst *Loader) _impl() keys.ServiceLoader {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

func NewLoader() *Loader {
	return new(Loader)
}
