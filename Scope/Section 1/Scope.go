package main

import "fmt"

/*variable x has package-level scope.*/
var x int = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}