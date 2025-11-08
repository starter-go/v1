package collections

import "github.com/starter-go/v1/lang"

type Setter interface {
	Error() error

	SetObject(name string, value any)

	SetInt(name string, value int)
	SetInt8(name string, value int8)
	SetInt16(name string, value int16)
	SetInt32(name string, value int32)
	SetInt64(name string, value int64)

	SetUInt(name string, value uint)
	SetUInt8(name string, value uint8)
	SetUInt16(name string, value uint16)
	SetUInt32(name string, value uint32)
	SetUInt64(name string, value uint64)

	SetFloat32(name string, value float32)
	SetFloat64(name string, value float64)

	SetBool(name string, value bool)
	SetByte(name string, value byte)
	SetRune(name string, value rune)
	SetString(name string, value string)

	SetComplex64(name string, value complex64)
	SetComplex128(name string, value complex128)

	SetHex(name string, value lang.Hex)
	SetBase64(name string, value lang.Base64)
	SetUUID(name string, value lang.UUID)
	SetTimeStamp(name string, value lang.TimeStamp)
	SetTimeSpan(name string, value lang.TimeSpan)
}
