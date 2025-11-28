package main

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"

	"github.com/starter-go/v1/platforms"
)

func main() {

	mod := platforms.GetThisModule()
	goos := runtime.GOOS
	goar := runtime.GOARCH

	// properties
	props := make(map[string]string)
	props["module.str"] = mod.String()

	props["module.name"] = mod.Name()
	props["module.rev"] = strconv.Itoa(mod.Revision())
	props["module.ver"] = mod.Version()

	props["runtime.goos"] = goos
	props["runtime.goarch"] = goar

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
