package libafs

import (
	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
)

////////////////////////////////////////////////////////////////////////////////

func createNewPlatformAPI() implementation.PlatformAPI {
	return new(innerLinuxPlatformAPI)
}

////////////////////////////////////////////////////////////////////////////////

type innerLinuxPlatformAPI struct{}

// ListRoots implements implementation.PlatformAPI.
func (i *innerLinuxPlatformAPI) ListRoots() []afs.Path {
	panic("unimplemented")
}

// NormalizePath implements implementation.PlatformAPI.
func (i *innerLinuxPlatformAPI) NormalizePath(path afs.Path) afs.Path {
	panic("unimplemented")
}
