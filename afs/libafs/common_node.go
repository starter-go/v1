package libafs

import (
	"os"
	"strings"

	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
)

////////////////////////////////////////////////////////////////////////////////

type innerCommonNodeBuilder struct {
	context      *implementation.Context
	rawPath      afs.Path
	targetPath   afs.Path
	pathElements afs.PathElementList
	simpleName   afs.PathElement
}

func (inst *innerCommonNodeBuilder) init(path afs.Path) {
	elist := path.Elements()
	inst.rawPath = path
	inst.pathElements = elist
}

func (inst *innerCommonNodeBuilder) onBuild() error {

	elements := inst.pathElements
	target, err := elements.Normalize()
	if err != nil {
		return err
	}

	inst.pathElements = target
	inst.targetPath = target.Path()
	inst.simpleName = inst.getSimpleName(target)

	return nil
}

func (inst *innerCommonNodeBuilder) getSimpleName(elements afs.PathElementList) afs.PathElement {
	count := len(elements)
	for i := count - 1; i >= 0; i-- {
		item := elements[i]
		if !item.IsEmpty() {
			return item
		}
	}
	return ""
}

func (inst *innerCommonNodeBuilder) build() afs.Node {

	err := inst.onBuild()
	if err != nil {
		return nil
	}

	node := new(innerCommonNode)
	node.context = inst.context
	node.simpleName = inst.simpleName
	node.path = inst.targetPath
	return node
}

////////////////////////////////////////////////////////////////////////////////

type innerCommonNode struct {
	context    *implementation.Context
	path       afs.Path
	simpleName afs.PathElement
	metaCache  afs.Meta
}

func (inst *innerCommonNode) _impl() (afs.Node, afs.File, afs.Directory, afs.Link) {
	return inst, inst, inst, inst
}

// GetNameSuffix implements afs.File.
func (inst *innerCommonNode) GetNameSuffix() string {
	str := inst.simpleName.String()
	idx := strings.LastIndexByte(str, '.')
	if idx < 0 {
		return ""
	}
	suffix := str[idx:]
	return suffix
}

// GetNameSuffixLower implements afs.File.
func (inst *innerCommonNode) GetNameSuffixLower() string {
	suffix := inst.GetNameSuffix()
	return strings.ToLower(suffix)
}

// GetNameSuffixUpper implements afs.File.
func (inst *innerCommonNode) GetNameSuffixUpper() string {
	suffix := inst.GetNameSuffix()
	return strings.ToUpper(suffix)
}

func (inst *innerCommonNode) innerGetMeta(reload bool) afs.Meta {

	meta := inst.metaCache

	if reload {
		meta = nil
	}

	if meta == nil {
		api := inst.context.PlatformAPI
		meta = api.LoadMeta(inst)
		inst.metaCache = meta
	}

	return meta
}

// GetSize implements afs.File.
func (inst *innerCommonNode) GetSize(reload bool) afs.FileSize {
	meta := inst.innerGetMeta(reload)
	return meta.Size()
}

// GetFileSystem implements afs.Node.
func (inst *innerCommonNode) GetFileSystem() afs.FileSystem {
	return inst.context.FS
}

// GetName implements afs.Node.
func (inst *innerCommonNode) GetName() string {
	return inst.simpleName.String()
}

// GetParent implements afs.Node.
func (inst *innerCommonNode) GetParent() afs.Directory {
	node := inst.GetHref("..")
	if node == nil {
		return nil
	}
	return node.(afs.Directory)
}

// GetPath implements afs.Node.
func (inst *innerCommonNode) GetPath() afs.Path {
	return inst.path
}

// GetURI implements afs.Node.
func (inst *innerCommonNode) GetURI() afs.URI {
	path := inst.path
	api := inst.context.FullAPI
	return api.PathToUri(path)
}

// String implements afs.Node.
func (inst *innerCommonNode) String() string {
	return inst.path.String()
}

// GetTarget implements afs.Link.
func (inst *innerCommonNode) GetTarget() afs.Node {

	//todo ..
	panic("unimplemented")
}

// GetIO implements afs.File.
func (inst *innerCommonNode) GetIO() afs.FileSystemIO {
	return inst.context.IO
}

////////////////////////////////////////////////////////////////////////////////
// Node interface implementation

func (inst *innerCommonNode) IsDir() bool {
	meta := inst.innerGetMeta(false)
	return meta.IsDir()
}

func (inst *innerCommonNode) IsFile() bool {
	meta := inst.innerGetMeta(false)
	return meta.IsFile()
}

func (inst *innerCommonNode) IsLink() bool {
	meta := inst.innerGetMeta(false)
	return meta.IsLink()
}

func (inst *innerCommonNode) Exists() bool {
	meta := inst.innerGetMeta(false)
	return meta.Exists()
}

func (inst *innerCommonNode) GetMeta(opt *afs.Options) afs.Meta {
	reload := false
	if opt != nil {
		reload = opt.Reload
	}
	return inst.innerGetMeta(reload)
}

////////////////////////////////////////////////////////////////////////////////
// Directory interface implementation

func (inst *innerCommonNode) ListNames() []string {

	name := inst.String()
	dst := make([]string, 0)
	src, err := os.ReadDir(name)

	if err == nil {
		for _, item := range src {
			simpleName := item.Name()
			dst = append(dst, simpleName)
		}
	}

	return dst
}

func (inst *innerCommonNode) ListPaths() []afs.Path {

	path1 := inst.String()
	dst := make([]afs.Path, 0)
	src, err := os.ReadDir(path1)

	if err == nil {
		api := inst.context.PlatformAPI
		for _, item := range src {
			simpleName := item.Name()
			path2 := path1 + "/" + simpleName
			path3, _, err := api.NormalizePath(afs.Path(path2))
			if err == nil {
				dst = append(dst, path3)
			}
		}
	}

	return dst
}

func (inst *innerCommonNode) ListNodes() []afs.Node {

	src := inst.ListPaths()
	dst := make([]afs.Node, 0)

	for _, it1 := range src {
		it2 := inst.context.FS.GetNode(it1)
		dst = append(dst, it2)
	}

	return dst
}

func (inst *innerCommonNode) GetChild(name afs.PathElement) afs.Node {
	str := string(name)
	return inst.GetHref(afs.Path("./" + str))
}

func (inst *innerCommonNode) GetHref(path afs.Path) afs.Node {

	pelist2 := path.Elements()
	if pelist2.IsAbsolute() {
		return inst.context.FS.GetNode(path)
	}

	// else,  pelist2.IsRelative()
	pelist1 := inst.path.Elements()
	pelist1n2 := append(pelist1, pelist2...)
	nextPathEL, err := pelist1n2.Normalize()
	if err != nil {
		return nil
	}
	nextPathStr := nextPathEL.Path()

	// build
	nb := new(innerCommonNodeBuilder)
	nb.context = inst.context
	nb.pathElements = nextPathEL
	nb.rawPath = nextPathStr
	nb.targetPath = nextPathStr

	return nb.build()
}

////////////////////////////////////////////////////////////////////////////////
// Link interface implementation

// Link interface doesn't require additional methods beyond Node interface
