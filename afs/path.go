package afs

import (
	"fmt"
	"os/user"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////
// path

// Path 以字符串的形式表示文件系统中的一个路径
type Path string

// 取路径的字符串形式
func (p Path) String() string {
	return string(p)
}

// 把路径字符串拆分成元素列表
func (p Path) Elements() PathElementList {
	const (
		sep  = '\n'
		sep1 = '\\'
		sep2 = '/'
	)
	str := p.String()

	str = strings.ReplaceAll(str, string(sep1), string(sep))
	str = strings.ReplaceAll(str, string(sep2), string(sep))

	src := strings.Split(str, string(sep))
	dst := make(PathElementList, 0)
	for _, item := range src {
		dst = append(dst, PathElement(item))
	}
	return dst
}

// 标准化路径
func (p Path) Normalize() (Path, error) {

	elist1 := p.Elements()
	elist2, err := elist1.Normalize()

	if err != nil {
		return "", err
	}

	p2 := elist2.Path()
	return p2, nil
}

////////////////////////////////////////////////////////////////////////////////
// path-element

type PathElement string

func (el PathElement) String() string {
	return string(el)
}

func (el PathElement) IsEmpty() bool {
	return (el == "")
}

func (el PathElement) IsDot() bool {
	return (el == ".")
}

func (el PathElement) IsDoubleDot() bool {
	return (el == "..")
}

func (el PathElement) IsUserHome() bool {
	return (el == "~")
}

////////////////////////////////////////////////////////////////////////////////
// path-element-list

// PathElements 表示一个路径元素列表.
// 重要: 对于相对路径, 必须 一律使用 '.' 作为开头元素
type PathElementList []PathElement

// 把元素组装成路径字符串
// 对于不同的系统平台,   统一按照 POSIX 规范的路径字符串处理,
// 如有特殊需要, 应该在后续步奏进行处理加工
func (list PathElementList) Path() Path {
	const sep = '/'
	// abs := list.IsAbsolute()
	builder := new(strings.Builder)
	for _, el := range list {
		if el.IsEmpty() {
			continue
		}
		builder.WriteRune(sep)
		builder.WriteString(el.String())
	}
	str := builder.String()
	return Path(str)
}

// 获取父元素列表
func (pe PathElementList) GetParent() (PathElementList, error) {
	size := len(pe)
	if size < 1 {
		return nil, fmt.Errorf("PathElementList: no parent")
	}
	parent := pe[0 : size-1]
	return parent, nil
}

// 获取指定名称的子元素列表
func (pe PathElementList) GetChild(name PathElement) PathElementList {
	return append(pe, name)
}

// 判断是否为绝对路径
func (list PathElementList) IsAbsolute() bool {
	return !list.IsRelative()
}

// 判断是否为相对路径
func (list PathElementList) IsRelative() bool {
	for _, el := range list {
		if el.IsEmpty() {
			continue
		} else if el.IsDot() || el.IsDoubleDot() {
			return true
		} else {
			return false
		}
	}
	return false
}

// 标准化路径元素列表
func (pe PathElementList) Normalize() (PathElementList, error) {

	src := pe
	tmp := make(PathElementList, 0)

	for _, el := range src {
		if el.IsEmpty() {
			continue
		} else if el.IsDot() {
			continue
		} else if el.IsDoubleDot() {
			parent, err := tmp.GetParent()
			if err != nil {
				return nil, err
			}
			tmp = parent
		} else if el.IsUserHome() {
			tmp = innerPathElementListGetUserHomeDir()
		} else {
			tmp = append(tmp, el)
		}
	}

	return tmp, nil
}

func innerPathElementListGetUserHomeDir() PathElementList {
	u, err := user.Current()
	if err != nil {
		// return make(PathElementList, 0)
		panic(err)
	}

	home := u.HomeDir
	if strings.ContainsRune(home, '~') {
		panic("bad home dir path: " + home)
	}

	path := Path(home)
	pelist1 := path.Elements()
	pelist2, err := pelist1.Normalize()
	if err != nil {
		panic(err)
	}

	return pelist2
}

////////////////////////////////////////////////////////////////////////////////
// EOF
