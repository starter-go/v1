package lang

import "github.com/starter-go/base/lang"

type Hex = lang.Hex

func HexFromString(str string) Hex {
	return Hex(str)
}

func HexFromBytes(b []byte) Hex {
	return lang.HexFromBytes(b)
}
