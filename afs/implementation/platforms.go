package implementation

import "github.com/starter-go/v1/afs"

// 表示系统平台特殊的接口
type PlatformAPI interface {
	NormalizePath(path afs.Path) (afs.Path, afs.PathElementList, error)

	NormalizePathEL(elements afs.PathElementList) (afs.Path, afs.PathElementList, error)

	LoadMeta(node afs.Node) afs.Meta

	ListRoots() []afs.Path
}

// 表示通用的接口
type CommonAPI interface {
	GetContext() *Context

	UriToPath(uri afs.URI) afs.Path

	PathToUri(path afs.Path) afs.URI
}

// 表示 afs.FS 全功能的接口
type FileSystemAPI interface {
	CommonAPI
	PlatformAPI
}
