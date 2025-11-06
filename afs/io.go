package afs

import "time"

// Options 包含几个通用的选项,对于不同的操作,某些字段可能是无效的
type Options struct {

	// 文件模式
	Mode FileMode

	// 读写标志位
	Flag int

	Owner UserID

	Group GroupID

	// 指示操作不要使用缓存
	Reload bool
}

////////////////////////////////////////////////////////////////////////////////

type FileSystemIO interface {
	FileSystem() FileSystem

	// file i/o

	ReadFile(file File, p *Options) ([]byte, error)

	WriteFile(file File, data []byte, p *Options) error

	// make dir

	Mkdir(dir Directory, p *Options) error

	Mkdirs(dir Directory, p *Options) error

	// change (mode|owner|group)

	Chmod(node Node, p *Options) error

	Chown(node Node, p *Options) error

	Chgrp(node Node, p *Options) error

	// set time

	SetCreatedAt(node Node, at time.Time) error

	SetUpdatedAt(node Node, at time.Time) error
}

////////////////////////////////////////////////////////////////////////////////

// NodeIO 为 Node 提供一组便捷的 IO 方法
type NodeIO interface {
	Node() Node

	ReadText(p *Options) (string, error)

	ReadBinary(p *Options) ([]byte, error)

	WriteText(text string, p *Options) error

	WriteBinary(data []byte, p *Options) error

	Mkdir(p *Options) error

	Mkdirs(p *Options) error
}

////////////////////////////////////////////////////////////////////////////////
