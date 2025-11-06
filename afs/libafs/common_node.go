package libafs

import (
	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
)

type innerCommonNode struct {
	context *implementation.Context
	path    afs.Path
}

func (inst *innerCommonNode) _impl() (afs.Node, afs.File, afs.Directory, afs.Link) {
	return inst, inst, inst, inst
}

////////////////////////////////////////////////////////////////////////////////
// Node interface implementation

func (inst *innerCommonNode) Name() string {
	// TODO: implement
	return ""
}

func (inst *innerCommonNode) Path() afs.Path {
	return inst.path
}

func (inst *innerCommonNode) URI() afs.URI {
	path := inst.path.String()
	return "file:///" + afs.URI(path)
}

func (inst *innerCommonNode) Parent() afs.Directory {
	// TODO: implement
	return nil
}

func (inst *innerCommonNode) FileSystem() afs.FileSystem {
	// TODO: implement
	return nil
}

func (inst *innerCommonNode) IsDir() bool {
	// TODO: implement
	return false
}

func (inst *innerCommonNode) IsFile() bool {
	// TODO: implement
	return false
}

func (inst *innerCommonNode) IsLink() bool {
	// TODO: implement
	return false
}

func (inst *innerCommonNode) Exists() bool {
	// TODO: implement
	return false
}

func (inst *innerCommonNode) GetMeta(opt *afs.Options) afs.Meta {
	// TODO: implement
	return nil
}

func (inst *innerCommonNode) IO() afs.NodeIO {
	// TODO: implement
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// File interface implementation

func (inst *innerCommonNode) Size() afs.FileSize {
	// TODO: implement
	return 0
}

func (inst *innerCommonNode) NameSuffix() string {
	// TODO: implement
	return ""
}

func (inst *innerCommonNode) NameSuffixLower() string {
	// TODO: implement
	return ""
}

func (inst *innerCommonNode) NameSuffixUpper() string {
	// TODO: implement
	return ""
}

////////////////////////////////////////////////////////////////////////////////
// Directory interface implementation

func (inst *innerCommonNode) ListNames() []string {
	// TODO: implement
	return nil
}

func (inst *innerCommonNode) ListPaths() []afs.Path {
	// TODO: implement
	return nil
}

func (inst *innerCommonNode) ListNodes() []afs.Node {
	// TODO: implement
	return nil
}

func (inst *innerCommonNode) GetChild(name afs.PathElement) afs.Node {
	// TODO: implement
	return nil
}

func (inst *innerCommonNode) GetHref(path afs.Path) afs.Node {
	// TODO: implement
	return nil
}

func (inst *innerCommonNode) Mkdir(opt *afs.Options) error {
	// TODO: implement
	return nil
}

func (inst *innerCommonNode) Mkdirs(opt *afs.Options) error {
	// TODO: implement
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Link interface implementation

// Link interface doesn't require additional methods beyond Node interface
