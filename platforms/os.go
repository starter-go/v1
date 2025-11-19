package platforms

import (
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

// OperatingSystemVersion 表示操作系统的版本. 对于不同的系统, 版本号的格式不尽相同. 此处不做过多限制
type OperatingSystemVersion string

// OperatingSystemType 表示操作系统的类型.
// 例如: 'linux'   , 'bsd' , 'windows' 等
type OperatingSystemType string

// OperatingSystemName 表示操作系统的名称 .
// 例如: 'ubuntu'| 'redhat' | 'macos'  等
type OperatingSystemName string

// OperatingSystemRevision 表示操作系统的修改版次, 它是一个整数, 以便于比较两个版本的先后次序
type OperatingSystemRevision int

// OST 是 OperatingSystemType 的简写形式
type OST = OperatingSystemType

////////////////////////////////////////////////////////////////////////////////

func (osv OperatingSystemVersion) String() string {
	return string(osv)
}

func (osv OperatingSystemVersion) Normalize() OperatingSystemVersion {

	str := osv.String()
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	return OperatingSystemVersion(str)

}

////////////////////////////////////////////////////////////////////////////////

func (osr OperatingSystemRevision) String() string {
	n := osr.Int()
	return strconv.Itoa(n)
}

func (osr OperatingSystemRevision) Int() int {
	return int(osr)
}

////////////////////////////////////////////////////////////////////////////////
// OST

func (ost OperatingSystemType) String() string {
	return string(ost)
}

// func (ost OperatingSystemType) Major() string {
// 	str := ost.String()
// 	idx := strings.IndexByte(str, '/')
// 	if idx < 0 {
// 		return str
// 	}
// 	return str[0:idx]
// }

// func (ost OperatingSystemType) Minor() string {
// 	str := ost.String()
// 	idx := strings.IndexByte(str, '/')
// 	if idx < 0 {
// 		return ""
// 	}
// 	return str[idx+1:]
// }

func (ost OperatingSystemType) Normalize() OperatingSystemType {

	str := ost.String()
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	return OperatingSystemType(str)
}

////////////////////////////////////////////////////////////////////////////////

func (name OperatingSystemName) String() string {
	return string(name)
}

func (name OperatingSystemName) Normalize() OperatingSystemName {

	str := name.String()
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	return OperatingSystemName(str)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
