package main

import _ "math/rand"

func main() {
	println("Hello World !")

	println("%v ", sum(4, 5))

	a, b := similarDatatypes(8, 9)

	println("", a, b)
}

/* NOTE: Unlike c ++ and dart in this language the datatype declaration succeeds the variable declaration.  and the return type preseeds the function lambda.  */

func sum(x int, y int) int {

	return x + y
}

/* NOTE: similoar functions can be declared all at once like this example below. */
func similarDatatypes(x, y int) (int, int) {

	return x, y
}
