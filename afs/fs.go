package afs

type FileSystem interface {
	GetNode(path Path) Node

	GetNodeWithURI(uri URI) Node

	ListRoots() []Node

	GetIO() FileSystemIO
}

// FS 是 FileSystem 的别名
type FS = FileSystem
