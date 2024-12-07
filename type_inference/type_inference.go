package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := 45.565465584

	b := reflect.TypeOf(a).Kind()

	fmt.Printf(" %q", b)
}
