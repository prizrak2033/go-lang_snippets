package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "vini, vedi, vici"
	fmt.Printf("strings.SplitAfterN(): %#v\n", strings.SplitAfterN(str, ", ", 2))
}
