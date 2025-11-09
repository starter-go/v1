package afs

import "os"

type FileMode = os.FileMode

type FileModeBuilder struct {
	mode FileMode
}

func (inst *FileModeBuilder) SetPerm(owner, group, other int) *FileModeBuilder {
	n1 := (0x07 & owner)
	n2 := (0x07 & group)
	n3 := (0x07 & other)
	mix := (n1 << 6) | (n2 << 3) | (n3)
	inst.mode = FileMode(mix)
	return inst
}

func (inst *FileModeBuilder) Mode() FileMode {
	return inst.mode
}
