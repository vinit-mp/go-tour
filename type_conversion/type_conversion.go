package main

import (
	"fmt"
	"reflect"
)

func main() {

	i := float64(42)

	f := int(i)

	d := uint8(f)
	ix := reflect.TypeOf(i).Kind()

	fx := reflect.TypeOf(f).Kind()

	dx := reflect.TypeOf(d).Kind()

	fmt.Printf("These are the values : %v %v %v", i, f, d)

	fmt.Printf("\n Run time type of variables : %q %q %q", ix, fx, dx)

}
