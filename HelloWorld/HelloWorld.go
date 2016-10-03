package main

/*Every Go program is made up of packages.*/

/*These are the import paths in use by the package. */
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*var i, j int = 1, 2
var c, python, java = true, false, "no!"
var (a = 5 b = 10 c = 1)

var c int

/*In go, you state the type after variable name.
When two or more consecutive named function parameters share a type,
you can omit the type from all but the last like so...
func add(x, y int) int { */

func add(x int, y int) int {
	return x + y
}

/*swap function*/
func swap(x, y string) (string, string) {
	return y, x
}

/*Go's return values may be named.
If so, they are treated as variables defined at the top of the function.

A return statement without arguments returns the named return values.
This is known as a "naked" return. Should only be used in short functions.*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c string = "Raffi"

func main() {
	fmt.Println(c)

	/*you can only use := inside a function as a short assignment statement
	instead of declaring the type i.e. var*/

	fmt.Printf("Hello World! My name is " + c + ".")
	fmt.Println("The time is", time.Now())
	fmt.Println(rand.Intn(10))

	/*factored import statement which takes the value and enters it in string.*/
	fmt.Printf("Now you have %g problems. ", math.Sqrt(7))

	/*Pi is an exported name*/
	fmt.Println(math.Pi)

	fmt.Println(add(10, 10))

	a := "hello"
	b := "hi"
	fmt.Println(swap(a, b))

	fmt.Println(split(20))

	/*i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128*/

	v := 1 + 2i // change me!
	fmt.Printf("v is of type %T\n", v)

	/*constants cannot use :=*/
	const Pi = 2.1
	fmt.Println(Pi)
}
