package common

import (
	"sort"
	"strings"
	"testing"

	"github.com/starter-go/v1/lang"
)

func TestPropertyTableCodec(t *testing.T) {

	table1 := make(map[string]string)
	codec := new(PropertyTableCodec)
	keys := []string{}

	table1["a"] = "11"
	table1["b"] = "22"
	table1["c"] = "33"
	table1["a.b"] = "44"
	table1["a.b.c"] = "55"
	table1["a.b.c.d"] = "66"

	data1, err := codec.Encode(table1)
	if err != nil {
		t.Fatal(err)
		return
	}

	table2, err := codec.Decode(data1, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	// t1
	t.Log("table1:")
	for k, v := range table1 {
		keys = append(keys, k)
		t.Logf("\t %s = %s", k, v)
	}

	//t2
	t.Log("table2:")
	for k, v := range table2 {
		keys = append(keys, k)
		t.Logf("\t %s = %s", k, v)
	}

	// t1+t2
	t.Log("table1+2:")
	for _, key := range keys {
		v1 := table1[key]
		v2 := table2[key]
		if v1 != v2 {
			t.Fatalf("bad key-value pair: key=[%s], value1=[%s], value2=[%s]", key, v1, v2)
			break
		}
		t.Logf("\t %s = %s", key, v2)
	}

	data2, err := codec.Encode(table2)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("data1 = {\n%s }", string(data1))
	t.Logf("data2 = {\n%s }", string(data2))

	t.Log("done")
}

func TestKeyValueGetterAndSetter(t *testing.T) {

	table := new(StringTable)
	mockBin := []byte{'M', 'o', 'c', 'k', '-', 'D', 'a', 't', 'a'}

	fnPrintKeyValuePairs := func(kv map[string]string) {
		keys := []string{}
		for k := range kv {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			v := kv[k]
			t.Logf("key-value[%d]:     %s = %s", i, k, v)
		}
	}

	fnPrintKeyValueForGetter := func(key string, fnGet func(key string) any) {
		value := fnGet(key)
		t.Logf("  getter.key_value_pair :  [%s] = [%v]", key, value)
	}

	// setter

	setter := table.Setter()

	setter.SetBool("k-bool", true)
	setter.SetByte("k-byte", 128)
	setter.SetComplex128("k-complex128", 3-4i)
	setter.SetComplex64("k-complex64", -1+2i)
	setter.SetFloat32("k-float32", 1.2345)
	setter.SetFloat64("k-float64", 9.876543)
	setter.SetRune("k-rune", 'x')
	setter.SetString("k-string", "hello")

	setter.SetInt("k-int", -1)
	setter.SetInt8("k-int8", -2)
	setter.SetInt16("k-int16", -3)
	setter.SetInt32("k-int32", -4)
	setter.SetInt64("k-int64", -5)

	setter.SetUint("k-uint", 1)
	setter.SetUint8("k-uint8", 2)
	setter.SetUint16("k-uint16", 3)
	setter.SetUint32("k-uint32", 4)
	setter.SetUint64("k-uint64", 5)

	setter.SetBase64("k-b64", lang.Base64FromBytes(mockBin))
	setter.SetHex("k-hex", lang.HexFromBytes(mockBin))
	setter.SetUUID("k-uuid", lang.UUIDFromString("mock-uuid"))
	setter.SetObject("k-object", new(strings.Builder))
	setter.SetTimeSpan("k-time-span", 5000)
	setter.SetTimeStamp("k-time-stamp", 1234567890)

	// list - all
	tmp := table.Export(nil)
	fnPrintKeyValuePairs(tmp)

	// getter

	getter := table.Getter()
	const def = 40

	fnPrintKeyValueForGetter("k-int", func(key string) any { return getter.GetInt(key, def) })
	fnPrintKeyValueForGetter("k-int8", func(key string) any { return getter.GetInt8(key, def) })
	fnPrintKeyValueForGetter("k-int16", func(key string) any { return getter.GetInt16(key, def) })
	fnPrintKeyValueForGetter("k-int32", func(key string) any { return getter.GetInt32(key, def) })
	fnPrintKeyValueForGetter("k-int64", func(key string) any { return getter.GetInt64(key, def) })

	fnPrintKeyValueForGetter("k-uint", func(key string) any { return getter.GetUint(key, def) })
	fnPrintKeyValueForGetter("k-uint8", func(key string) any { return getter.GetUint8(key, def) })
	fnPrintKeyValueForGetter("k-uint16", func(key string) any { return getter.GetUint16(key, def) })
	fnPrintKeyValueForGetter("k-uint32", func(key string) any { return getter.GetUint32(key, def) })
	fnPrintKeyValueForGetter("k-uint64", func(key string) any { return getter.GetUint64(key, def) })

	fnPrintKeyValueForGetter("k-bool", func(key string) any { return getter.GetBool(key, true, false) })
	fnPrintKeyValueForGetter("k-byte", func(key string) any { return getter.GetByte(key, '?') })
	fnPrintKeyValueForGetter("k-complex128", func(key string) any { return getter.GetComplex128(key, (1 + 1i)) })
	fnPrintKeyValueForGetter("k-complex64", func(key string) any { return getter.GetComplex64(key, (-2 - 2i)) })
	fnPrintKeyValueForGetter("k-float32", func(key string) any { return getter.GetFloat32(key, 3.14159) })
	fnPrintKeyValueForGetter("k-float64", func(key string) any { return getter.GetFloat64(key, 9.87654321) })
	fnPrintKeyValueForGetter("k-rune", func(key string) any { return getter.GetRune(key, '#') })
	fnPrintKeyValueForGetter("k-string", func(key string) any { return getter.GetString(key, "abcd") })

	fnPrintKeyValueForGetter("k-uuid", func(key string) any { return getter.GetUUID(key, "ffff") })

}
