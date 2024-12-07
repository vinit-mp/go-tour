//If a variable is not initialized to a specific value, then the initial value of the data type in given as follows

package main

import (
	"fmt"
)

func main() {
	var t int
	var i string
	var ts float32

	fmt.Printf("%v %q %v", t, i, ts)

}
