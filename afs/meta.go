package afs

import (
	"os"
	"time"
)

type FileSize int64

type FileMode = os.FileMode

type UserID int32

type GroupID int32

////////////////////////////////////////////////////////////////////////////////

type Meta interface {
	Node() Node

	CreatedAt() time.Time

	UpdatedAt() time.Time

	AccessedAt() time.Time

	Mode() FileMode

	Size() FileSize

	Owner() UserID

	Group() GroupID

	IsDir() bool

	IsFile() bool

	IsLink() bool

	Exists() bool
}

////////////////////////////////////////////////////////////////////////////////

type MetaBuilder struct {
	Node Node

	CreatedAt  time.Time
	UpdatedAt  time.Time
	AccessedAt time.Time

	Mode FileMode

	Size FileSize

	Owner UserID
	Group GroupID

	IsDir  bool
	IsFile bool
	IsLink bool
	Exists bool
}

func (inst *MetaBuilder) Build() Meta {

	now := time.Now()
	m := new(innerCommonMeta)

	m.node = inst.Node
	m.cachedAt = now

	m.createdAt = inst.CreatedAt
	m.updatedAt = inst.UpdatedAt
	m.accessedAt = inst.AccessedAt

	m.isFile = inst.IsFile
	m.isDir = inst.IsDir
	m.isLink = inst.IsLink
	m.exists = inst.Exists

	m.mode = inst.Mode
	m.size = inst.Size
	m.owner = inst.Owner
	m.group = inst.Group

	return m
}

////////////////////////////////////////////////////////////////////////////////

type innerCommonMeta struct {
	cachedAt time.Time

	node Node

	size FileSize
	mode FileMode

	createdAt  time.Time
	updatedAt  time.Time
	accessedAt time.Time

	owner UserID
	group GroupID

	isDir  bool
	isFile bool
	isLink bool
	exists bool
}

// AccessedAt implements Meta.
func (i *innerCommonMeta) AccessedAt() time.Time {
	return i.accessedAt
}

// CreatedAt implements Meta.
func (i *innerCommonMeta) CreatedAt() time.Time {
	return i.createdAt
}

// Exists implements Meta.
func (i *innerCommonMeta) Exists() bool {
	return i.exists
}

// Group implements Meta.
func (i *innerCommonMeta) Group() GroupID {
	return i.group
}

// IsDir implements Meta.
func (i *innerCommonMeta) IsDir() bool {
	return i.isDir
}

// IsFile implements Meta.
func (i *innerCommonMeta) IsFile() bool {
	return i.isFile
}

// IsLink implements Meta.
func (i *innerCommonMeta) IsLink() bool {
	return i.isLink
}

// Mode implements Meta.
func (i *innerCommonMeta) Mode() FileMode {
	return i.mode
}

// Node implements Meta.
func (i *innerCommonMeta) Node() Node {
	return i.node
}

// Owner implements Meta.
func (i *innerCommonMeta) Owner() UserID {
	return i.owner
}

// Size implements Meta.
func (i *innerCommonMeta) Size() FileSize {
	return i.size
}

// UpdatedAt implements Meta.
func (i *innerCommonMeta) UpdatedAt() time.Time {
	return i.updatedAt
}

////////////////////////////////////////////////////////////////////////////////
