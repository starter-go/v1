package main

import (
	"fmt"

	"github.com/starter-go/v1/platforms"
)

func main() {

	// module - info
	mod := platforms.GetThisModule()
	str := mod.String()

	fmt.Println("module: ", str)

	// platform - info
	pl := platforms.Current()
	str = pl.String()

	fmt.Println("platform: ", str)

}
