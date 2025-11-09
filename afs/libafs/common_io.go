package libafs

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
	"github.com/starter-go/v1/lang/ios"
)

type innerCommonFSIO struct {
	context *implementation.Context
}

// Chgrp implements afs.FileSystemIO.
func (inst *innerCommonFSIO) Chgrp(node afs.Node, p *afs.Options) error {
	// alias of Chown
	return inst.Chown(node, p)
}

// Chmod implements afs.FileSystemIO.
func (inst *innerCommonFSIO) Chmod(node afs.Node, opt *afs.Options) error {
	path := node.GetPath().String()
	mode := inst.innerGetMode(opt)
	return os.Chmod(path, mode)
}

// Chown implements afs.FileSystemIO.
func (inst *innerCommonFSIO) Chown(node afs.Node, p *afs.Options) error {
	path := node.GetPath().String()
	uid := inst.innerGetUID(p)
	gid := inst.innerGetGID(p)
	return os.Chown(path, uid, gid)
}

// Mkdir implements afs.FileSystemIO.
func (inst *innerCommonFSIO) Mkdir(dir afs.Directory, opt *afs.Options) error {
	path := dir.GetPath().String()
	mode := inst.innerGetMode(opt)
	return os.Mkdir(path, mode)
}

func (inst *innerCommonFSIO) innerGetFlag(p *afs.Options) int {
	if p == nil {
		return os.O_RDONLY
	}
	return p.Flag
}

func (inst *innerCommonFSIO) innerGetMode(p *afs.Options) os.FileMode {
	if p == nil {
		return 0644
	}
	return p.Mode
}

func (inst *innerCommonFSIO) innerGetUID(p *afs.Options) int {
	if p == nil {
		return -1
	}
	return int(p.User)
}

func (inst *innerCommonFSIO) innerGetGID(p *afs.Options) int {
	if p == nil {
		return -1
	}
	return int(p.Group)
}

// Mkdirs implements afs.FileSystemIO.
func (inst *innerCommonFSIO) Mkdirs(dir afs.Directory, opt *afs.Options) error {
	path := dir.GetPath().String()
	mode := inst.innerGetMode(opt)
	return os.MkdirAll(path, mode)
}

// OpenReader implements afs.FileSystemIO.
func (inst *innerCommonFSIO) OpenReader(file afs.File, opt *afs.Options) (io.ReadCloser, error) {
	path := file.GetPath().String()
	mode := inst.innerGetMode(opt)
	flag := inst.innerGetFlag(opt)
	return os.OpenFile(path, flag, mode)
}

// OpenWriter implements afs.FileSystemIO.
func (inst *innerCommonFSIO) OpenWriter(file afs.File, opt *afs.Options) (io.WriteCloser, error) {
	path := file.GetPath().String()
	mode := inst.innerGetMode(opt)
	flag := inst.innerGetFlag(opt)
	return os.OpenFile(path, flag, mode)
}

// ReadBinary implements afs.FileSystemIO.
func (inst *innerCommonFSIO) ReadBinary(file afs.File, opt *afs.Options) ([]byte, error) {
	return inst.ReadFile(file, opt)
}

// ReadFile implements afs.FileSystemIO.
func (inst *innerCommonFSIO) ReadFile(file afs.File, opt *afs.Options) ([]byte, error) {
	reader, err := inst.OpenReader(file, opt)
	if err != nil {
		return nil, err
	}
	defer ios.Close(reader)
	return io.ReadAll(reader)
}

// ReadText implements afs.FileSystemIO.
func (inst *innerCommonFSIO) ReadText(file afs.File, opt *afs.Options) (string, error) {
	data, err := inst.ReadFile(file, opt)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// SetCreatedAt implements afs.FileSystemIO.
func (inst *innerCommonFSIO) SetCreatedAt(node afs.Node, t time.Time) error {
	api := inst.context.PlatformAPI
	return api.SetNodeCreatedAt(node, t)
}

// SetUpdatedAt implements afs.FileSystemIO.
func (inst *innerCommonFSIO) SetUpdatedAt(node afs.Node, mt time.Time) error {
	path := node.GetPath().String()
	return os.Chtimes(path, mt, mt)
}

// WriteBinary implements afs.FileSystemIO.
func (inst *innerCommonFSIO) WriteBinary(file afs.File, data []byte, p *afs.Options) error {
	return inst.WriteFile(file, data, p)
}

// WriteFile implements afs.FileSystemIO.
func (inst *innerCommonFSIO) WriteFile(file afs.File, data []byte, opt *afs.Options) error {

	if opt == nil {

		fmb := new(afs.FileModeBuilder)
		fmb.SetPerm(6, 4, 4)

		opt = new(afs.Options)
		opt.Flag = os.O_TRUNC | os.O_CREATE | os.O_WRONLY
		opt.Mode = fmb.Mode()
		opt.User = -1
		opt.Group = -1
	}

	wtr, err := inst.OpenWriter(file, opt)
	if err != nil {
		return err
	}
	defer ios.Close(wtr)

	have, err := wtr.Write(data)
	if err != nil {
		return err
	}

	want := len(data)
	if want != have {
		return fmt.Errorf("FileSystemIO.WriteFile: bad count of bytes, want:%d, have:%d", want, have)
	}

	return nil
}

// WriteText implements afs.FileSystemIO.
func (inst *innerCommonFSIO) WriteText(file afs.File, text string, p *afs.Options) error {
	data := []byte(text)
	return inst.WriteFile(file, data, p)
}

func (inst *innerCommonFSIO) _impl() afs.FileSystemIO {
	return inst
}
