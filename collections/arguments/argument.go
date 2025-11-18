package arguments

import (
	"fmt"
	"strings"
)

type Argument struct {
	rawValue string
	value    string
	index    int
	next     *Argument
	list     *List
	used     bool
}

func (inst *Argument) Use() {
	inst.used = true
}

func (inst *Argument) IsUsed() bool {
	return inst.used
}

func (inst *Argument) Exists() bool {
	return (inst != nil)
}

func (inst *Argument) IsFlag() bool {
	if inst == nil {
		return false
	}
	value := inst.value
	return strings.HasPrefix(value, "-")
}

func (inst *Argument) GetNext() *Argument {
	return inst.next
}

func (inst *Argument) GetList() *List {
	return inst.list
}

func (inst *Argument) TryGetKeyValue() (key string, value string, err error) {

	str := inst.value
	idx := strings.IndexByte(str, '=')
	if idx < 1 {
		return "", "", fmt.Errorf("")
	}

	p1 := str[0:idx]
	p2 := str[idx+1:]
	p1 = strings.TrimSpace(p1)
	p2 = strings.TrimSpace(p2)

	return p1, p2, nil
}

func (inst *Argument) HasMore() bool {
	if inst == nil {
		return false
	}
	if inst.next == nil {
		return false
	}
	return true
}

func (inst *Argument) String() string {
	return inst.value
}

func (inst *Argument) Value() string {
	return inst.value
}

func (inst *Argument) GetPureValue() string {

	const (
		mark1   = string('\'')
		mark2   = string('"')
		minSize = 2
	)

	value := inst.value
	size := len(value)
	if size >= minSize {
		hasMark1 := strings.HasPrefix(value, mark1) && strings.HasSuffix(value, mark1)
		hasMark2 := strings.HasPrefix(value, mark2) && strings.HasSuffix(value, mark2)
		if hasMark1 || hasMark2 {
			return value[1 : size-1]
		}
	}

	return value
}

func (inst *Argument) Index() int {
	return inst.index
}
