package afs

import (
	"io"
	"time"
)

type NodeIO interface {

	// change (mode|owner|group)

	Chmod(node Node, p *Options) error

	Chown(node Node, p *Options) error

	Chgrp(node Node, p *Options) error

	// set time

	SetCreatedAt(node Node, at time.Time) error

	SetUpdatedAt(node Node, at time.Time) error
}

type DirectoryIO interface {

	// make dir

	Mkdir(dir Directory, p *Options) error

	Mkdirs(dir Directory, p *Options) error
}

type FileIO interface {

	// file i/o

	ReadFile(file File, p *Options) ([]byte, error)

	WriteFile(file File, data []byte, p *Options) error

	// extends

	ReadText(file File, p *Options) (string, error)

	ReadBinary(file File, p *Options) ([]byte, error)

	WriteText(file File, text string, p *Options) error

	WriteBinary(file File, data []byte, p *Options) error

	OpenReader(file File, p *Options) (io.ReadCloser, error)

	OpenWriter(file File, p *Options) (io.WriteCloser, error)
}

type LinkIO interface {
}

type FileSystemIO interface {
	DirectoryIO
	FileIO
	LinkIO
	NodeIO
}

////////////////////////////////////////////////////////////////////////////////
