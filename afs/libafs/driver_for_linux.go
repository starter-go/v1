package libafs

import (
	"os"
	"syscall"
	"time"

	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
)

////////////////////////////////////////////////////////////////////////////////

func createNewPlatformAPI() implementation.PlatformAPI {
	return new(innerLinuxPlatformAPI)
}

////////////////////////////////////////////////////////////////////////////////

type innerLinuxPlatformAPI struct{}

// ListRoots implements implementation.PlatformAPI.
func (inst *innerLinuxPlatformAPI) ListRoots() []afs.Path {
	root := afs.Path("/")
	return []afs.Path{root}
}

func (inst *innerLinuxPlatformAPI) innerConvertTime(src syscall.Timespec) time.Time {
	return time.Unix(src.Sec, src.Nsec)
}

// LoadMeta implements implementation.PlatformAPI.
func (inst *innerLinuxPlatformAPI) LoadMeta(node afs.Node) afs.Meta {

	path := node.GetPath()
	dst := new(afs.MetaBuilder)

	src, err := os.Stat(path.String())
	if err != nil {
		return dst.Build()
	}

	mode := src.Mode()
	size := src.Size()
	isDir := src.IsDir()
	modTime := src.ModTime()
	more := src.Sys()

	dst.Node = node
	dst.Size = afs.FileSize(size)
	dst.Mode = mode
	dst.UpdatedAt = modTime
	dst.IsDir = isDir
	dst.IsFile = !isDir
	dst.IsLink = false // todo ...
	dst.Exists = true

	st, ok := more.(*syscall.Stat_t)
	if ok {
		dst.Owner = afs.UserID(st.Uid)
		dst.Group = afs.GroupID(st.Gid)
		dst.AccessedAt = inst.innerConvertTime(st.Atim)
		dst.CreatedAt = inst.innerConvertTime(st.Ctim)
	}

	return dst.Build()
}

// NormalizePath implements implementation.PlatformAPI.
func (inst *innerLinuxPlatformAPI) NormalizePath(path afs.Path) (afs.Path, afs.PathElementList, error) {
	elist := path.Elements()
	return inst.innerNormalizePathEL(elist)
}

// NormalizePathEL implements implementation.PlatformAPI.
func (inst *innerLinuxPlatformAPI) NormalizePathEL(elements afs.PathElementList) (afs.Path, afs.PathElementList, error) {
	return inst.innerNormalizePathEL(elements)
}

func (inst *innerLinuxPlatformAPI) innerNormalizePathEL(elements afs.PathElementList) (afs.Path, afs.PathElementList, error) {

	abs := elements.IsAbsolute()

	elist, err := elements.Normalize()
	if err != nil {
		return "", nil, err
	}

	path := elist.Path()
	if abs {
		if path == "" {
			path = "/"
		}
	}

	return path, elist, nil
}

func (inst *innerLinuxPlatformAPI) _impl() implementation.PlatformAPI {
	return inst
}
