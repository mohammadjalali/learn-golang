package main

import (
	"fmt"
)

func f(a []int) {
	a[1] = 100
}

func main() {
	a := make([]int, 2, 100)
	f(a)
	fmt.Println(a)
}
