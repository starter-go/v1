package afs

import (
	"os"
	"time"
)

type FileSize uint64

type FileMode = os.FileMode

type UserID int

type GroupID int

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
