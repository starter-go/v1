package common

import (
	"fmt"
	"maps"
	"sort"
	"strconv"

	"github.com/starter-go/v1/collections"
	"github.com/starter-go/v1/lang"
	"github.com/starter-go/v1/lang/threads"
)

type StringTable struct {
	table    map[string]string
	strategy threads.Strategy
	locker   threads.Locker
	prec     int // 表示浮点数的精度
}

func (inst *StringTable) innerMakeTable() map[string]string {
	return make(map[string]string)
}

func (inst *StringTable) innerGetTable() map[string]string {
	table := inst.table
	if table == nil {
		table = inst.innerMakeTable()
		inst.table = table
	}
	return table
}

func (inst *StringTable) innerGetLocker() threads.Locker {
	return inst.locker
}

func (inst *StringTable) Init(strategy threads.Strategy) {
	if strategy == nil {
		strategy = threads.GetStrategy(threads.Fast)
	}
	inst.table = make(map[string]string)
	inst.locker = strategy.NewLocker()
	inst.strategy = strategy
	inst.prec = 9
}

func (inst *StringTable) Reset() {
	inst.table = make(map[string]string)
}

func (inst *StringTable) Getter() collections.Getter {

	getter := &innerStringTableGetter{
		strTable: inst,
	}
	return getter
}

func (inst *StringTable) Setter() collections.Setter {

	setter := &innerStringTableSetter{
		strTable: inst,
	}
	return setter
}

func (inst *StringTable) Keys(needSort bool) []string {

	locker := inst.innerGetLocker()
	locker.Lock()
	defer locker.Unlock()

	table := inst.innerGetTable()
	keys := make([]string, 0)
	for key := range table {
		keys = append(keys, key)
	}
	if needSort {
		sort.Strings(keys)
	}
	return keys
}

func (inst *StringTable) Export(dst map[string]string) map[string]string {

	locker := inst.innerGetLocker()
	locker.Lock()
	defer locker.Unlock()

	src := inst.table
	if dst == nil {
		dst = make(map[string]string)
	}
	if src != nil {
		maps.Copy(dst, src)
	}
	return dst
}

func (inst *StringTable) Import(src map[string]string) {

	if src == nil {
		return
	}

	locker := inst.innerGetLocker()
	locker.Lock()
	defer locker.Unlock()

	dst := inst.innerGetTable()
	maps.Copy(dst, src)
}

func (inst *StringTable) Set(key, value string) {

	locker := inst.innerGetLocker()
	locker.Lock()
	defer locker.Unlock()

	table := inst.innerGetTable()
	table[key] = value
}

func (inst *StringTable) Get(key string) string {

	locker := inst.innerGetLocker()
	locker.Lock()
	defer locker.Unlock()

	table := inst.table
	if table == nil {
		return ""
	}
	return table[key]
}

////////////////////////////////////////////////////////////////////////////////
// impl: setter

type innerStringTableSetter struct {
	strTable *StringTable
	err      error
}

// GetPrecision implements collections.Setter.
func (inst *innerStringTableSetter) GetPrecision() int {

	return inst.strTable.prec
}

// SetPrecision implements collections.Setter.
func (inst *innerStringTableSetter) SetPrecision(prec int) {

	inst.strTable.prec = prec
}

// Error implements collections.Setter.
func (inst *innerStringTableSetter) Error() error {
	return inst.err
}

// SetBase64 implements collections.Setter.
func (inst *innerStringTableSetter) SetBase64(name string, value lang.Base64) {
	str := value.String()
	inst.strTable.Set(name, str)
}

// SetBool implements collections.Setter.
func (inst *innerStringTableSetter) SetBool(name string, value bool) {
	str := strconv.FormatBool(value)
	inst.strTable.Set(name, str)
}

// SetByte implements collections.Setter.
func (inst *innerStringTableSetter) SetByte(name string, value byte) {
	str := strconv.Itoa(int(value))
	inst.strTable.Set(name, str)
}

// SetComplex128 implements collections.Setter.
func (inst *innerStringTableSetter) SetComplex128(name string, value complex128) {

	prec := inst.strTable.prec
	str := strconv.FormatComplex(value, 'f', prec, 128)
	inst.strTable.Set(name, str)
}

// SetComplex64 implements collections.Setter.
func (inst *innerStringTableSetter) SetComplex64(name string, value complex64) {

	prec := inst.strTable.prec
	str := strconv.FormatComplex(complex128(value), 'f', prec, 64)
	inst.strTable.Set(name, str)
}

// SetFloat32 implements collections.Setter.
func (inst *innerStringTableSetter) SetFloat32(name string, value float32) {

	prec := inst.strTable.prec
	str := strconv.FormatFloat(float64(value), 'f', prec, 32)
	inst.strTable.Set(name, str)
}

// SetFloat64 implements collections.Setter.
func (inst *innerStringTableSetter) SetFloat64(name string, value float64) {

	prec := inst.strTable.prec
	str := strconv.FormatFloat(float64(value), 'f', prec, 64)
	inst.strTable.Set(name, str)
}

// SetHex implements collections.Setter.
func (inst *innerStringTableSetter) SetHex(name string, value lang.Hex) {

	str := value.String()
	inst.strTable.Set(name, str)

}

// SetInt implements collections.Setter.
func (inst *innerStringTableSetter) SetInt(name string, value int) {

	str := strconv.FormatInt(int64(value), 10)
	inst.strTable.Set(name, str)

}

// SetInt16 implements collections.Setter.
func (inst *innerStringTableSetter) SetInt16(name string, value int16) {

	str := strconv.FormatInt(int64(value), 10)
	inst.strTable.Set(name, str)

}

// SetInt32 implements collections.Setter.
func (inst *innerStringTableSetter) SetInt32(name string, value int32) {

	str := strconv.FormatInt(int64(value), 10)
	inst.strTable.Set(name, str)

}

// SetInt64 implements collections.Setter.
func (inst *innerStringTableSetter) SetInt64(name string, value int64) {

	str := strconv.FormatInt(value, 10)
	inst.strTable.Set(name, str)

}

// SetInt8 implements collections.Setter.
func (inst *innerStringTableSetter) SetInt8(name string, value int8) {

	str := strconv.FormatInt(int64(value), 10)
	inst.strTable.Set(name, str)

}

// SetObject implements collections.Setter.
func (inst *innerStringTableSetter) SetObject(name string, value any) {

	str := ""
	ss, ok := value.(fmt.Stringer)
	if ok {
		str = ss.String()
	}
	inst.strTable.Set(name, str)

}

// SetRune implements collections.Setter.
func (inst *innerStringTableSetter) SetRune(name string, value rune) {

	var buffer [1]rune
	buffer[0] = value
	str := string(buffer[:])
	inst.strTable.Set(name, str)

}

// SetString implements collections.Setter.
func (inst *innerStringTableSetter) SetString(name string, value string) {

	inst.strTable.Set(name, value)

}

// SetTimeSpan implements collections.Setter.
func (inst *innerStringTableSetter) SetTimeSpan(name string, value lang.TimeSpan) {

	num := int64(value)
	inst.SetInt64(name, num)

}

// SetTimeStamp implements collections.Setter.
func (inst *innerStringTableSetter) SetTimeStamp(name string, value lang.TimeStamp) {

	num := value.Int()
	inst.SetInt64(name, num)

}

// SetUInt implements collections.Setter.
func (inst *innerStringTableSetter) SetUint(name string, value uint) {

	str := strconv.FormatUint(uint64(value), 10)
	inst.strTable.Set(name, str)

}

// SetUInt16 implements collections.Setter.
func (inst *innerStringTableSetter) SetUint16(name string, value uint16) {

	str := strconv.FormatUint(uint64(value), 10)
	inst.strTable.Set(name, str)

}

// SetUInt32 implements collections.Setter.
func (inst *innerStringTableSetter) SetUint32(name string, value uint32) {

	str := strconv.FormatUint(uint64(value), 10)
	inst.strTable.Set(name, str)

}

// SetUInt64 implements collections.Setter.
func (inst *innerStringTableSetter) SetUint64(name string, value uint64) {

	str := strconv.FormatUint(uint64(value), 10)
	inst.strTable.Set(name, str)

}

// SetUInt8 implements collections.Setter.
func (inst *innerStringTableSetter) SetUint8(name string, value uint8) {

	str := strconv.FormatUint(uint64(value), 10)
	inst.strTable.Set(name, str)

}

// SetUUID implements collections.Setter.
func (inst *innerStringTableSetter) SetUUID(name string, value lang.UUID) {

	str := value.String()
	inst.SetString(name, str)

}

// func (inst *innerStringTableSetter) init() {}

////////////////////////////////////////////////////////////////////////////////
// impl: getter

type innerStringTableGetter struct {
	strTable *StringTable
	err      error
	required bool
}

func (inst *innerStringTableGetter) handleError(err error) {

	if err == nil {
		return
	}

	if inst.required {
		inst.err = err
	}

}

func (inst *innerStringTableGetter) innerGetValue(key string) string {

	value := inst.strTable.Get(key)
	if inst.required {
		if value == "" {
			err := fmt.Errorf("innerStringTableGetter: no value with name [%s]", key)
			inst.handleError(err)
		}
	}
	return value
}

// Error implements collections.Getter.
func (inst *innerStringTableGetter) Error() error {
	return inst.err
}

// GetBase64 implements collections.Getter.
func (inst *innerStringTableGetter) GetBase64(name string, valueDefault ...lang.Base64) lang.Base64 {

	str := inst.innerGetValue(name)
	if str != "" {
		b64a := lang.Base64FromString(str)
		b64b := b64a.Bytes()
		return lang.Base64FromBytes(b64b)
	}
	for _, item := range valueDefault {
		return item
	}
	return ""

}

// GetBool implements collections.Getter.
func (inst *innerStringTableGetter) GetBool(name string, valueDefault ...bool) bool {

	str := inst.innerGetValue(name)
	if str != "" {
		value, err := strconv.ParseBool(str)
		if err == nil {
			return value
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return false

}

// GetByte implements collections.Getter.
func (inst *innerStringTableGetter) GetByte(name string, valueDefault ...byte) byte {

	str := inst.innerGetValue(name)
	if str != "" {
		b := str[0]
		return b
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetComplex128 implements collections.Getter.
func (inst *innerStringTableGetter) GetComplex128(name string, valueDefault ...complex128) complex128 {

	const (
		bitSize = 128
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseComplex(str, bitSize)
		if err == nil {
			return num
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetComplex64 implements collections.Getter.
func (inst *innerStringTableGetter) GetComplex64(name string, valueDefault ...complex64) complex64 {

	const (
		bitSize = 64
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseComplex(str, bitSize)
		if err == nil {
			return complex64(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetFloat32 implements collections.Getter.
func (inst *innerStringTableGetter) GetFloat32(name string, valueDefault ...float32) float32 {

	const (
		bitSize = 32
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseFloat(str, bitSize)
		if err == nil {
			return float32(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetFloat64 implements collections.Getter.
func (inst *innerStringTableGetter) GetFloat64(name string, valueDefault ...float64) float64 {

	const (
		bitSize = 64
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseFloat(str, bitSize)
		if err == nil {
			return num
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetHex implements collections.Getter.
func (inst *innerStringTableGetter) GetHex(name string, valueDefault ...lang.Hex) lang.Hex {

	str := inst.innerGetValue(name)
	if str != "" {
		hx1 := lang.HexFromString(str)
		hx2 := hx1.Bytes()
		return lang.HexFromBytes(hx2)
	}
	for _, item := range valueDefault {
		return item
	}
	return ""

}

// GetInt implements collections.Getter.
func (inst *innerStringTableGetter) GetInt(name string, valueDefault ...int) int {

	const (
		base    = 10
		bitSize = 0
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseInt(str, base, bitSize)
		if err == nil {
			return int(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetInt16 implements collections.Getter.
func (inst *innerStringTableGetter) GetInt16(name string, valueDefault ...int16) int16 {

	const (
		base    = 10
		bitSize = 16
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseInt(str, base, bitSize)
		if err == nil {
			return int16(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetInt32 implements collections.Getter.
func (inst *innerStringTableGetter) GetInt32(name string, valueDefault ...int32) int32 {

	const (
		base    = 10
		bitSize = 32
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseInt(str, base, bitSize)
		if err == nil {
			return int32(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetInt64 implements collections.Getter.
func (inst *innerStringTableGetter) GetInt64(name string, valueDefault ...int64) int64 {

	const (
		base    = 10
		bitSize = 64
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseInt(str, base, bitSize)
		if err == nil {
			return int64(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetInt8 implements collections.Getter.
func (inst *innerStringTableGetter) GetInt8(name string, valueDefault ...int8) int8 {

	const (
		base    = 10
		bitSize = 8
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseInt(str, base, bitSize)
		if err == nil {
			return int8(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetObject implements collections.Getter.
func (inst *innerStringTableGetter) GetObject(name string, valueDefault ...any) any {

	str := inst.innerGetValue(name)
	if str != "" {
		return str
	}
	for _, item := range valueDefault {
		return item
	}
	return ""

}

// GetRune implements collections.Getter.
func (inst *innerStringTableGetter) GetRune(name string, valueDefault ...rune) rune {

	const (
		base    = 10
		bitSize = 64
	)

	str := inst.innerGetValue(name)
	if str != "" {
		b := str[0]
		return rune(b)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetString implements collections.Getter.
func (inst *innerStringTableGetter) GetString(name string, valueDefault ...string) string {

	value := inst.innerGetValue(name)

	if value == "" {
		if len(valueDefault) > 0 {
			return valueDefault[0]
		}
	}
	return value
}

// GetTimeSpan implements collections.Getter.
func (inst *innerStringTableGetter) GetTimeSpan(name string, valueDefault ...lang.TimeSpan) lang.TimeSpan {

	const (
		base    = 10
		bitSize = 64
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseInt(str, base, bitSize)
		if err == nil {
			return lang.TimeSpan(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetTimeStamp implements collections.Getter.
func (inst *innerStringTableGetter) GetTimeStamp(name string, valueDefault ...lang.TimeStamp) lang.TimeStamp {

	const (
		base    = 10
		bitSize = 64
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseInt(str, base, bitSize)
		if err == nil {
			return lang.TimeStamp(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetUUID implements collections.Getter.
func (inst *innerStringTableGetter) GetUUID(name string, valueDefault ...lang.UUID) lang.UUID {

	str := inst.innerGetValue(name)
	if str != "" {
		uuid := lang.UUIDFromString(str)
		return uuid.Normalize()
	}
	for _, item := range valueDefault {
		return item
	}
	uuid := lang.UUIDFromString("0000")
	return uuid.Normalize()

}

// GetUint implements collections.Getter.
func (inst *innerStringTableGetter) GetUint(name string, valueDefault ...uint) uint {

	const (
		base    = 10
		bitSize = 0
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseUint(str, base, bitSize)
		if err == nil {
			return uint(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetUint16 implements collections.Getter.
func (inst *innerStringTableGetter) GetUint16(name string, valueDefault ...uint16) uint16 {

	const (
		base    = 10
		bitSize = 16
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseUint(str, base, bitSize)
		if err == nil {
			return uint16(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetUint32 implements collections.Getter.
func (inst *innerStringTableGetter) GetUint32(name string, valueDefault ...uint32) uint32 {

	const (
		base    = 10
		bitSize = 32
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseUint(str, base, bitSize)
		if err == nil {
			return uint32(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetUint64 implements collections.Getter.
func (inst *innerStringTableGetter) GetUint64(name string, valueDefault ...uint64) uint64 {

	const (
		base    = 10
		bitSize = 64
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseUint(str, base, bitSize)
		if err == nil {
			return uint64(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// GetUint8 implements collections.Getter.
func (inst *innerStringTableGetter) GetUint8(name string, valueDefault ...uint8) uint8 {

	const (
		base    = 10
		bitSize = 8
	)

	str := inst.innerGetValue(name)
	if str != "" {
		num, err := strconv.ParseUint(str, base, bitSize)
		if err == nil {
			return uint8(num)
		}
		inst.handleError(err)
	}
	for _, item := range valueDefault {
		return item
	}
	return 0

}

// Optional implements collections.Getter.
func (inst *innerStringTableGetter) Optional() collections.Getter {
	inst.required = false
	return inst
}

// Required implements collections.Getter.
func (inst *innerStringTableGetter) Required() collections.Getter {
	inst.required = true
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
