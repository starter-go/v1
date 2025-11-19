package platforms

import "strings"

// ArchName 表示指令集架构的名称. 例如: 'x86', 'amd64', 'armv7'
type ArchName string

func (an ArchName) String() string {
	return string(an)
}

func (an ArchName) Normalize() ArchName {

	str := an.String()
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	return ArchName(str)
}
