package main
import "fmt"

func main(){
	// ● assigns an int to a variable
	a := 99

	// ● prints that int in decimal, binary, and hex
	fmt.Printf("%d\t%b\t%x", a, a, a)

	// ● shifts the bits of that int over 1 position to the left, and assigns that to a variable
	b := a << 1
	fmt.Printf("%d\t%b\t%x", b, b, b)
}




// Write a program that
// ● assigns an int to a variable
// ● prints that int in decimal, binary, and hex
// ● shifts the bits of that int over 1 position to the left, and assigns that to a variable
// ● prints that variable in decimal, binary, and hex
// solution: https://play.golang.org/p/Ms964T8SbH