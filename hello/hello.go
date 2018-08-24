package main

import "fmt"
import "git/go-tests/goMath"

func main() {
	fmt.Println("Hello World")
	fmt.Println(goMath.Add(2,5))
	fmt.Println(goMath.Sub(2,5))
	fmt.Println(goMath.Mult(2,5))
	fmt.Println(goMath.Div(2,5))
}
