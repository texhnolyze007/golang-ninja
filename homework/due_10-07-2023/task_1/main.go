package main

import "fmt"

func fibonacci() func() int {
	a := -1
	b := 1
	return func() int {
		a += b
		a, b = b, a
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
