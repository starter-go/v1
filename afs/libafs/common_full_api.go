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

// LoadMeta implements implementation.FileSystemAPI.
func (inst *innerCommonFullAPIImpl) LoadMeta(node afs.Node) afs.Meta {

	// 委托给  platform-api

	panic("unimplemented")
}

// NormalizePath implements implementation.FileSystemAPI.
func (inst *innerCommonFullAPIImpl) NormalizePath(path afs.Path) (afs.Path, afs.PathElementList, error) {

	// 委托给  platform-api

	panic("unimplemented")
}

// NormalizePathEL implements implementation.FileSystemAPI.
func (inst *innerCommonFullAPIImpl) NormalizePathEL(elements afs.PathElementList) (afs.Path, afs.PathElementList, error) {

	// 委托给  platform-api

	panic("unimplemented")
}

func (inst *innerCommonFullAPIImpl) GetContext() *implementation.Context {
	return inst.context
}

func (inst *innerCommonFullAPIImpl) ListRoots() []afs.Path {

	// 委托给  platform-api

	//todo ...
	return nil
}

func (inst *innerCommonFullAPIImpl) PathToUri(path afs.Path) afs.URI {
	//todo ...
	return ""
}

func (inst *innerCommonFullAPIImpl) UriToPath(uri afs.URI) afs.Path {
	//todo ...
	return ""
}
