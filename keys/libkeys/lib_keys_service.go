package libkeys

import "github.com/starter-go/v1/keys"

type innerLibKeysService struct {
}

// GetDriverManager implements keys.Service.
func (inst *innerLibKeysService) GetDriverManager() keys.DriverManager {
	panic("unimplemented")
}

// GetDriverRegistry implements keys.Service.
func (inst *innerLibKeysService) GetDriverRegistry() keys.DriverRegistry {
	panic("unimplemented")
}

func (inst *innerLibKeysService) _impl() keys.Service {
	return inst
}
