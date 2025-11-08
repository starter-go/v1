package libafs

import (
	"io"
	"time"

	"github.com/starter-go/v1/afs"
)

type innerCommonFSIO struct {
}

// Chgrp implements afs.FileSystemIO.
func (inst *innerCommonFSIO) Chgrp(node afs.Node, p *afs.Options) error {
	panic("unimplemented")
}

// Chmod implements afs.FileSystemIO.
func (inst *innerCommonFSIO) Chmod(node afs.Node, p *afs.Options) error {
	panic("unimplemented")
}

// Chown implements afs.FileSystemIO.
func (inst *innerCommonFSIO) Chown(node afs.Node, p *afs.Options) error {
	panic("unimplemented")
}

// Mkdir implements afs.FileSystemIO.
func (inst *innerCommonFSIO) Mkdir(dir afs.Directory, p *afs.Options) error {
	panic("unimplemented")
}

// Mkdirs implements afs.FileSystemIO.
func (inst *innerCommonFSIO) Mkdirs(dir afs.Directory, p *afs.Options) error {
	panic("unimplemented")
}

// OpenReader implements afs.FileSystemIO.
func (inst *innerCommonFSIO) OpenReader(file afs.File, p *afs.Options) (io.ReadCloser, error) {
	panic("unimplemented")
}

// OpenWriter implements afs.FileSystemIO.
func (inst *innerCommonFSIO) OpenWriter(file afs.File, p *afs.Options) (io.WriteCloser, error) {
	panic("unimplemented")
}

// ReadBinary implements afs.FileSystemIO.
func (inst *innerCommonFSIO) ReadBinary(file afs.File, p *afs.Options) ([]byte, error) {
	panic("unimplemented")
}

// ReadFile implements afs.FileSystemIO.
func (inst *innerCommonFSIO) ReadFile(file afs.File, p *afs.Options) ([]byte, error) {
	panic("unimplemented")
}

// ReadText implements afs.FileSystemIO.
func (inst *innerCommonFSIO) ReadText(file afs.File, p *afs.Options) (string, error) {
	panic("unimplemented")
}

// SetCreatedAt implements afs.FileSystemIO.
func (inst *innerCommonFSIO) SetCreatedAt(node afs.Node, at time.Time) error {
	panic("unimplemented")
}

// SetUpdatedAt implements afs.FileSystemIO.
func (inst *innerCommonFSIO) SetUpdatedAt(node afs.Node, at time.Time) error {
	panic("unimplemented")
}

// WriteBinary implements afs.FileSystemIO.
func (inst *innerCommonFSIO) WriteBinary(file afs.File, data []byte, p *afs.Options) error {
	panic("unimplemented")
}

// WriteFile implements afs.FileSystemIO.
func (inst *innerCommonFSIO) WriteFile(file afs.File, data []byte, p *afs.Options) error {
	panic("unimplemented")
}

// WriteText implements afs.FileSystemIO.
func (inst *innerCommonFSIO) WriteText(file afs.File, text string, p *afs.Options) error {
	panic("unimplemented")
}

func (inst *innerCommonFSIO) _impl() afs.FileSystemIO {
	return inst
}
