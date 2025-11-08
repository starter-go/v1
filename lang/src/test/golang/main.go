package main

import (
	"fmt"

	"github.com/starter-go/v1/lang"
)

func main() {
	m := lang.GetThisModule()
	str := m.String()
	fmt.Println(str + "#test")
}
