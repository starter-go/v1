package libafs

import (
	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
)

////////////////////////////////////////////////////////////////////////////////

func createNewPlatformAPI() implementation.PlatformAPI {
	return new(innerWindowsPlatformAPI)
}

////////////////////////////////////////////////////////////////////////////////

type innerWindowsPlatformAPI struct{}

// ListRoots implements implementation.PlatformAPI.
func (i *innerWindowsPlatformAPI) ListRoots() []afs.Path {
	panic("unimplemented")
}

// NormalizePath implements implementation.PlatformAPI.
func (i *innerWindowsPlatformAPI) NormalizePath(path afs.Path) afs.Path {
	panic("unimplemented")
}
