package buckets

import (
	"strconv"
	"strings"
)

const (
	theModuleName     = "github.com/starter-go/v1/buckets"
	theModuleVersion  = "buckets/v0.9.0"
	theModuleRevision = 0
)

func GetModuleInfo() string {

	rev := strconv.Itoa(theModuleRevision)
	b := new(strings.Builder)

	b.WriteString("module:/info?updated_at=2025-11-06")
	b.WriteString("&name=" + theModuleName)
	b.WriteString("&ver=" + theModuleVersion)
	b.WriteString("&rev=" + rev)

	return b.String()
}
