package main

import (
	"fmt"
	"sort"

	"github.com/starter-go/v1/lang"
	"github.com/starter-go/v1/platforms"
)

func main() {

	mod := platforms.GetThisModule()
	pl := platforms.Current()
	str1 := mod.String()
	str2 := pl.String()
	now := lang.Now()

	fmt.Println(str1)
	fmt.Println(str2)

	// properties
	props := pl.GetProperties(nil)
	props["test-module"] = mod.String()
	props["test-platform"] = pl.String()
	props["test-date"] = now.String()

	keys := make([]string, 0)
	for k := range props {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("properties:")
	for _, name := range keys {
		value := props[name]
		fmt.Println("\t", name, "=", value)
	}

}
