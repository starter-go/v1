package libafs

import (
	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
)

type innerCommonFullAPIImpl struct {
	context *implementation.Context
}

func (inst *innerCommonFullAPIImpl) _impl() implementation.FileSystemAPI {
	return inst
}

func (inst *innerCommonFullAPIImpl) GetContext() *implementation.Context {
	return inst.context
}

func (inst *innerCommonFullAPIImpl) ListRoots() []afs.Path {

	//todo ...
	return nil
}

func (inst *innerCommonFullAPIImpl) NormalizePath(path afs.Path) afs.Path {
	//todo ...
	return ""
}

func (inst *innerCommonFullAPIImpl) PathToUri(path afs.Path) afs.URI {
	//todo ...
	return ""
}

func (inst *innerCommonFullAPIImpl) UriToPath(uri afs.URI) afs.Path {
	//todo ...
	return ""
}
