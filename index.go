package main

import _ "math/rand"

func main() {
	println("Hello World !")

	println("%v ", sum(4, 5))

}

/* NOTE: Unlike c ++ and dart in this language the datatype declaration succeeds the variable declaration.  and the return type preseeds the function lambda.  */

func sum(x int, y int) int {

	return x + y
}
