package main

import (
	"heima_extend"
	"fmt"
)

func main() {
	a := 10
	b := 20
	heima_extend.Swap(a, b)
	fmt.Printf("a=%d,b=%d\n", a, b)
	heima_extend.Swapx(&a, &b)
	fmt.Printf("a=%d,b=%d\n", a, b)
}
