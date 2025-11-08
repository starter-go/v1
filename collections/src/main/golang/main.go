package main

import (
	"fmt"

	"github.com/starter-go/v1/collections"
)

func main() {

	m := collections.GetThisModule()
	s := m.String()
	fmt.Println(s + "#main")

}
