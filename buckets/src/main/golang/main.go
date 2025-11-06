package main

import (
	"fmt"

	"github.com/starter-go/v1/buckets"
)

func main() {
	info := buckets.GetModuleInfo()
	fmt.Println(info)
}
