package afs

type Driver interface {
	GetFS() FS

	CreateNewFS() FS
}
