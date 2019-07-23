package main

import "fmt"

// for init; condition; post {}

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		} else {
			fmt.Println(n)
		}
	}

	for i := 0; i <= 1; i++ {
		fmt.Println("E:")
		for j := 0; j <= 2; j++ {
			fmt.Println(" EEEEEEEEEEE ")
			for k := 0; k <= 4; k++ {
				if j == 2 {
					break
				}
				fmt.Println(" EEE ")
			}
		}
	}

	// F
	for i := 0; i <= 1; i++ {
		fmt.Println("F:")
		for j := 0; j <= 1; j++ {
			fmt.Println("FFFFFFFFFFFF")
			for k := 0; k <= 1; k++ {
				fmt.Println("FFF")
			}
		}
		fmt.Println("           ")
	}

	// S
	for i := 0; i <= 1; i++ {
		fmt.Println("S:")
		fmt.Println("SSSSSSSSSS")
		for j := 0; j <= 2; j++ {
			fmt.Println("SS")
		}
		fmt.Println("SSSSSSSSSS")
		fmt.Println("SSSSSSSSSS")
		for j := 0; j <= 2; j++ {
			fmt.Println("        SS")
		}
		fmt.Println("SSSSSSSSSS")
		fmt.Println("           ")
	}

	// S using three for loops
	// draw two S's
	for s := 0; s <= 1; s++ {
		fmt.Println("S using three for loops:")
		//draw three full rows
		for i := 0; i <= 2; i++ {
			fmt.Println("SSSSSSSSSS")
			// draw the columns
			for j := 0; j <= 1; j++ {
				if i == 0 {
					fmt.Println("SS")
				} else if i == 1 {
					fmt.Println("        SS")
				} else {
					break
				}
			}
		}
	}

	// print outer and inner loops, with space between outer loops
	for o := 0; o <= 10; o++ {
		fmt.Println("")
		fmt.Printf("The outer loop: %d\n", o)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t\t The inner loop: %d\n", i)
		}
	}

	// countdown
	for c := 10; c > 0; c-- {
		fmt.Println(c)
	}
	fmt.Println("Blast off!")

	// countdown version II
	c := 10
	for c > 0 {
		fmt.Println(c)
		c--
	}
	fmt.Println("Engines Ignite!")

	// countdown version III - an eternal loop with a break!
	c = 10
	for {
		if c > 1 {
			fmt.Println(c)
		}
		c--
		fmt.Println("Key Ignition!")
	}
}
