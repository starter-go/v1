package lang

import "github.com/starter-go/base/lang"

type Base64 = lang.Base64

func Base64FromString(str string) Base64 {
	return Base64(str)
}

func Base64FromBytes(b []byte) Base64 {
	return lang.Base64FromBytes(b)
}
