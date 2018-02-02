package main

import (
	"fmt"
	"math/rand"
)

func insertion(arg []int) []int {
	for _ = range arg {
		for i := 0; i < len(arg)-1; i += 2 {
			if arg[i] > arg[i+1] {
				arg[i], arg[i+1] = arg[i+1], arg[i]
			}
		}
	}
	return arg
}

func main() {
	slice := make([]int, 0)
	for i := 0; i <= 100; i++ {
		slice = append(slice, rand.Int()/10000000000000)
	}
	fmt.Println(slice)
	fmt.Println(insertion(slice))
	fmt.Println("Θ(n)")
	fmt.Println("O(n)")

}
