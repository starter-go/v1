package collections

import "github.com/starter-go/v1/lang"

type Getter interface {
	Error() error

	Optional() Getter

	Required() Getter

	GetObject(name string, valueDefault ...any) any

	GetInt(name string, valueDefault ...int) int
	GetInt8(name string, valueDefault ...int8) int8
	GetInt16(name string, valueDefault ...int16) int16
	GetInt32(name string, valueDefault ...int32) int32
	GetInt64(name string, valueDefault ...int64) int64

	GetUint(name string, valueDefault ...uint) uint
	GetUint8(name string, valueDefault ...uint8) uint8
	GetUint16(name string, valueDefault ...uint16) uint16
	GetUint32(name string, valueDefault ...uint32) uint32
	GetUint64(name string, valueDefault ...uint64) uint64

	GetFloat32(name string, valueDefault ...float32) float32
	GetFloat64(name string, valueDefault ...float64) float64

	GetString(name string, valueDefault ...string) string
	GetByte(name string, valueDefault ...byte) byte
	GetRune(name string, valueDefault ...rune) rune
	GetBool(name string, valueDefault ...bool) bool

	GetComplex64(name string, valueDefault ...complex64) complex64
	GetComplex128(name string, valueDefault ...complex128) complex128

	GetHex(name string, valueDefault ...lang.Hex) lang.Hex
	GetBase64(name string, valueDefault ...lang.Base64) lang.Base64
	GetUUID(name string, valueDefault ...lang.UUID) lang.UUID
	GetTimeStamp(name string, valueDefault ...lang.TimeStamp) lang.TimeStamp
	GetTimeSpan(name string, valueDefault ...lang.TimeSpan) lang.TimeSpan
}
