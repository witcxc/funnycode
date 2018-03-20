package main

import (
	"def"
	"fmt"
	"test"
)

func main() {
	fmt.Printf("\nhello[%d]", def.Con)
	def.Con += 2
	test.Change()
	fmt.Printf("\nhello[%d]", def.Con)

}
