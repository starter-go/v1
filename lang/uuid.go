package lang

import "github.com/starter-go/base/lang"

type UUID = lang.UUID

func UUIDFromString(str string) UUID {
	return UUID(str)
}
