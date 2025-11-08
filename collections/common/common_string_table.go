package common

import "github.com/starter-go/v1/collections"

type StringTable struct {
	table map[string]string
}

func (inst *StringTable) innerGetTable() map[string]string {}

func (inst *StringTable) Getter() collections.Getter {}

func (inst *StringTable) Setter() collections.Setter {}

func (inst *StringTable) Keys() []string {}

func (inst *StringTable) Export() map[string]string {}

func (inst *StringTable) Import() map[string]string {}
