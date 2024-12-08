package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 6, 9, 5, 6, 1, 6, 14, 6, 52, 9, 5, 5, 8}

	for i := 0; i < len(a); i++ {
		fmt.Printf("%v", a[i])

	}

}
