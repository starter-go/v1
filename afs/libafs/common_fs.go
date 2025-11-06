package libafs

import (
	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
)

type innerCommonFS struct {
	context *implementation.Context
}

func (inst *innerCommonFS) _impl() afs.FS {
	return inst
}

func (inst *innerCommonFS) GetNode(path afs.Path) afs.Node {

	api := inst.context.FullAPI
	path = api.NormalizePath(path)

	node := new(innerCommonNode)
	node.context = inst.context
	node.path = path
	return node
}

func (inst *innerCommonFS) GetNodeWithURI(uri afs.URI) afs.Node {

	// const prefix = "file:/"
	api := inst.context.FullAPI
	path := api.UriToPath(uri)
	return inst.GetNode(path)
}

func (inst *innerCommonFS) ListRoots() []afs.Node {

	api := inst.context.FullAPI
	src := api.ListRoots()
	dst := make([]afs.Node, 0)

	for _, item1 := range src {
		item2 := inst.GetNode(item1)
		dst = append(dst, item2)
	}

	return dst
}

func (inst *innerCommonFS) IO() afs.FileSystemIO {
	return inst.context.IO
}
