// Rune(int32) to string conversion

package main

import (
	"fmt"
)

func main() {

	var r rune = 'a'

	var b string = string(r)
	fmt.Printf("%v, %T", b, b)

}
