package main

import (
	"fmt"
	"heima_extend"
)

//type People struct {
//	name string // 这里没有逗号
//	sex  byte
//	age  int
//}

func main() {
	p := heima_extend.People{"Bob", 'm', 27};
	fmt.Println(p.Age);
}
