package libafs

import (
	"strings"
	"time"

	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
)

type innerCommonFullAPIImpl struct {
	context *implementation.Context
}

// SetNodeCreatedAt implements implementation.FileSystemAPI.
func (inst *innerCommonFullAPIImpl) SetNodeCreatedAt(node afs.Node, t time.Time) error {
	// 委托给  platform-api
	api := inst.context.PlatformAPI
	return api.SetNodeCreatedAt(node, t)
}

func (inst *innerCommonFullAPIImpl) _impl() implementation.FileSystemAPI {
	return inst
}

// LoadMeta implements implementation.FileSystemAPI.
func (inst *innerCommonFullAPIImpl) LoadMeta(node afs.Node) afs.Meta {
	// 委托给  platform-api
	api := inst.context.PlatformAPI
	return api.LoadMeta(node)
}

// NormalizePath implements implementation.FileSystemAPI.
func (inst *innerCommonFullAPIImpl) NormalizePath(path afs.Path) (afs.Path, afs.PathElementList, error) {
	// 委托给  platform-api
	api := inst.context.PlatformAPI
	return api.NormalizePath(path)
}

// NormalizePathEL implements implementation.FileSystemAPI.
func (inst *innerCommonFullAPIImpl) NormalizePathEL(elements afs.PathElementList) (afs.Path, afs.PathElementList, error) {
	// 委托给  platform-api
	api := inst.context.PlatformAPI
	return api.NormalizePathEL(elements)
}

func (inst *innerCommonFullAPIImpl) GetContext() *implementation.Context {
	return inst.context
}

func (inst *innerCommonFullAPIImpl) ListRoots() []afs.Path {
	// 委托给  platform-api
	api := inst.context.PlatformAPI
	return api.ListRoots()
}

func (inst *innerCommonFullAPIImpl) PathToUri(path afs.Path) afs.URI {

	src := path.Elements()
	builder := new(strings.Builder)
	builder.WriteString("file://")
	count := 0

	src, err := src.Normalize()
	if err != nil {
		builder.WriteString("/log/.error?message=")
		builder.WriteString(err.Error())
		str := builder.String()
		return afs.URI(str)
	}

	for _, el := range src {
		builder.WriteRune('/')
		builder.WriteString(el.String())
		count++
	}

	if count == 0 {
		builder.WriteRune('/')
	}

	str := builder.String()
	return afs.URI(str)
}

func (inst *innerCommonFullAPIImpl) UriToPath(uri afs.URI) afs.Path {

	const prefix = "file:/"
	str := uri.String()
	str = strings.TrimPrefix(str, prefix)

	path := afs.Path(str)
	elist := path.Elements()

	path2, _, err2 := inst.NormalizePathEL(elist)
	if err2 != nil {
		return afs.ErrorPath
	}

	return path2
}
