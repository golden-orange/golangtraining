package main
import "fmt"


type beer int
var x beer

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}




// Hands-on exercise #4
// ● FYI - nice documentation and new terminology “ underlying type ”
// ○ https://golang.org/ref/spec#Types
// For this exercise
// 1. Create your own type. Have the underlying type be an int.
// 2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR”
// keyword
// 3. in func main
// a. print out the value of the variable “x”
// b. print out the type of the variable “x”
// c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
// d. print out the value of the variable “x”
// code: here’s the solution: https://play.golang.org/p/snm4WuuYmG