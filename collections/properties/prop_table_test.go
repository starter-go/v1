package properties

import (
	"crypto/sha1"
	"fmt"
	"strings"
	"testing"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v1/lang/threads"
)

func TestPropertyTableFormatterAndParser(t *testing.T) {

	tester := new(innerPropertyTableTester)
	tab1 := tester.prepare()

	text, err := Format(tab1)
	if err != nil {
		t.Error(err)
		return
	}

	tab2, err := Parse(text, nil)
	if err != nil {
		t.Error(err)
		return
	}

	msg1 := tester.formatAll(tab1)
	msg2 := tester.formatAll(tab2)
	msg3 := text

	t.Log("table-1 = \n", msg1)
	t.Log("table-2 = \n", msg2)
	t.Log("formatted.text = \n", msg3)
}

func TestPropertyTableGetterAndSetter(t *testing.T) {

	tester := new(innerPropertyTableTester)
	tab := tester.prepare()

	now := lang.Now()
	bin := sha1.Sum([]byte(now.String()))
	b64 := lang.Base64FromBytes(bin[:])
	hex := lang.HexFromBytes(bin[:])
	uuid := lang.ParseUUID("1111-2222-3333-4444")

	const (
		keyUint = "key.uint"
		keyInt  = "key.int"

		keyFloat32 = "key.float32"
		keyFloat64 = "key.float64"

		keyComplex64  = "key.complex64"
		keyComplex128 = "key.complex128"

		keyBool   = "key.bool"
		keyString = "key.string"
		keyByte   = "key.byte"
		keyRune   = "key.rune"

		keyHex       = "key.hex"
		keyBase64    = "key.base64"
		keyUUID      = "key.uuid"
		keyTimeStamp = "key.time-stamp"
		keyTimeSpan  = "key.time-span"
	)

	// setter

	setter := tab.Setter()
	setter.SetPrecision(20)

	setter.SetInt(keyInt, -10)
	setter.SetUint(keyUint, 10)
	setter.SetFloat32(keyFloat32, 3.14)
	setter.SetFloat64(keyFloat64, 0.123456789012345678)
	setter.SetComplex64(keyComplex64, (1.2345 + 2i))
	setter.SetComplex128(keyComplex128, (-3 - 4.2333i))

	setter.SetBool(keyBool, true)
	setter.SetByte(keyByte, '#')
	setter.SetRune(keyRune, '*')
	setter.SetString(keyString, "hello,table")

	setter.SetHex(keyHex, hex)
	setter.SetBase64(keyBase64, b64)
	setter.SetUUID(keyUUID, uuid)
	setter.SetTimeSpan(keyTimeSpan, -6000)
	setter.SetTimeStamp(keyTimeStamp, now)

	// getter

	getter := tab.Getter()

	getter.GetBool("")

	msg1 := tester.formatAll(tab)
	t.Log("table = \n", msg1)
}

func TestPropertyTableImportAndExport(t *testing.T) {

	tester := new(innerPropertyTableTester)
	tab1 := tester.prepare()
	tab2 := NewTable()
	tab3 := NewTable()

	tmp1 := make(map[string]string)
	tmp2 := make(map[string]string)
	tmp3 := make(map[string]string)

	tmp1 = tab1.Export(tmp1)

	tab2.Import(tmp1)
	tmp2 = tab2.Export(tmp2)

	for k, v := range tmp2 {
		k = strings.ToLower(k)
		tmp3[k] = v
	}

	tab3.Import(tmp3)

	msg1 := tester.formatAll(tab1)
	msg2 := tester.formatAll(tab2)
	msg3 := tester.formatAll(tab3)

	t.Log("table1 = \n", msg1)
	t.Log("table2 = \n", msg2)
	t.Log("table3 = \n", msg3)
}

func TestPropertyTableKeys(t *testing.T) {

	tester := new(innerPropertyTableTester)
	tab := tester.prepare()

	keys := tab.Keys()

	t.Log("keys:")
	for _, k := range keys {
		t.Log(k)
	}
}

func TestPropertyTableReset(t *testing.T) {

	tester := new(innerPropertyTableTester)
	tab := tester.prepare()

	str1 := tester.formatAll(tab)
	tab.Reset()
	str2 := tester.formatAll(tab)

	t.Log("reset().before: \n", str1)
	t.Log("reset().after: \n", str2)
}

////////////////////////////////////////////////////////////////////////////////

type innerPropertyTableTester struct {
}

func (inst *innerPropertyTableTester) prepare() Table {

	src := make(map[string]string)
	dst := NewTableWithMode(threads.Safe)

	src["fruit.Apple.name"] = "苹果"
	src["fruit.Apple.icon"] = "🍎"

	src["fruit.Banana.name"] = "香蕉"
	src["fruit.Banana.icon"] = "🍌"

	src["fruit.Orange.name"] = "橙子"
	src["fruit.Orange.icon"] = "🍊"

	src["fruit.Grape.name"] = "葡萄"
	src["fruit.Grape.icon"] = "🍇"

	src["fruit.Strawberry.name"] = "草莓"
	src["fruit.Strawberry.icon"] = "🍓"

	src["fruit.Watermelon.name"] = "西瓜"
	src["fruit.Watermelon.icon"] = "🍉"

	src["fruit.Peach.name"] = "桃子"
	src["fruit.Peach.icon"] = "🍑"

	src["fruit.Pear.name"] = "梨"
	src["fruit.Pear.icon"] = "🍐"

	src["fruit.Pineapple.name"] = "菠萝"
	src["fruit.Pineapple.icon"] = "🍍"

	src["fruit.Mango.name"] = "芒果"
	src["fruit.Mango.icon"] = "🥭"

	src["fruit.Cherry.name"] = "樱桃"
	src["fruit.Cherry.icon"] = "🍒"

	src["fruit.Lemon.name"] = "柠檬"
	src["fruit.Lemon.icon"] = "🍋"

	src["fruit.Coconut.name"] = "椰子"
	src["fruit.Coconut.icon"] = "🥥"

	src["fruit.Kiwi.name"] = "猕猴桃"
	src["fruit.Kiwi.icon"] = "🥝"

	src["fruit.Blueberry.name"] = "蓝莓"
	src["fruit.Blueberry.icon"] = "🫐"

	dst.Import(src)
	return dst
}

func (inst *innerPropertyTableTester) formatAll(tab Table) string {

	builder := new(strings.Builder)
	keys := tab.Keys()

	for i, k := range keys {
		v := tab.GetProperty(k)
		str := fmt.Sprintf("property[%d]: %s = %s", i, k, v)
		builder.WriteString(str)
		builder.WriteRune('\n')
	}

	return builder.String()
}
