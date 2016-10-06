package main


import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true

	/*Places the value of the variable
	in the string.*/
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println()

	/*Prints the type of the variable.*/
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
