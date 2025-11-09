package afs

type Node interface {

	// attributes

	String() string

	GetName() string

	GetPath() Path

	GetURI() URI

	GetParent() Directory

	ListParents() []Directory

	CountParents() int

	GetFileSystem() FileSystem

	IsDir() bool

	IsFile() bool

	IsLink() bool

	Exists() bool

	// 如果 reload=1, 表示必须重新加载元信息，并把之前缓存的 meta 扔掉
	GetMeta(opt *Options) Meta

	GetIO() FileSystemIO
}

////////////////////////////////////////////////////////////////////////////////

type File interface {
	Node

	GetSize(reload bool) FileSize

	// suffix

	GetNameSuffix() string

	GetNameSuffixLower() string

	GetNameSuffixUpper() string
}

////////////////////////////////////////////////////////////////////////////////

type Directory interface {
	Node

	// list children:

	ListNames() []string

	ListPaths() []Path

	ListNodes() []Node

	// ref

	GetChild(name PathElement) Node

	GetHref(path Path) Node
}

////////////////////////////////////////////////////////////////////////////////

type Link interface {
	Node

	GetTarget() Node
}

////////////////////////////////////////////////////////////////////////////////
