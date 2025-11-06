package afs

type Node interface {
	Name() string

	Path() Path

	URI() URI

	Parent() Directory

	FileSystem() FileSystem

	IsDir() bool
	IsFile() bool
	IsLink() bool

	Exists() bool

	// 如果 reload=1, 表示必须重新加载元信息，并把之前缓存的 meta 扔掉
	GetMeta(opt *Options) Meta

	IO() NodeIO
}

////////////////////////////////////////////////////////////////////////////////

type File interface {
	Node

	Size() FileSize

	NameSuffix() string
	NameSuffixLower() string
	NameSuffixUpper() string
}

////////////////////////////////////////////////////////////////////////////////

type Directory interface {
	Node

	// list children:

	ListNames() []string
	ListPaths() []Path
	ListNodes() []Node

	GetChild(name PathElement) Node
	GetHref(path Path) Node

	Mkdir(opt *Options) error
	Mkdirs(opt *Options) error
}

////////////////////////////////////////////////////////////////////////////////

type Link interface {
	Node
}

////////////////////////////////////////////////////////////////////////////////
