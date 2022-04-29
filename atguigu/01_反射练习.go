package main

/*
 * @Description:
 * @Author: Liuhongq
 * @Date: 2022-04-13 05:12:00
 * @LastEditTime: 2022-04-21 05:18:40
 * @LastEditors: Liuhongq
 * @Reference:
 */

import (
	"fmt"
	"reflect"
)

/*
1)编写一个Cal结构体，有两个字段Num1,和Num2。
2)方法GetSub(name string)
3)使用反射遍历Cal结构体所有的字段信息.
4)使用反射机制完成对GetSub的调用，输出形式为
"tom完成了减法运行，8-3= 5"
*/
type Cal struct {
	Num1 int
	Num2 int
}

func (this Cal) GetSub(name string) {
	fmt.Printf("%s完成了减法运行，%d - %d = %d\n", name, this.Num1, this.Num2, this.Num1-this.Num2)
}
func main() {
	var (
		model *Cal
		rt    reflect.Type
		rv    reflect.Value
	)
	rt = reflect.TypeOf(model)
	rt = rt.Elem()
	rv = reflect.New(rt)
	model = rv.Interface().(*Cal)
	// model = &Cal{}
	// rv = reflect.ValueOf(model)
	rv = rv.Elem()

	rv.FieldByName("Num1").SetInt(8)
	rv.FieldByName("Num2").SetInt(3)

	fmt.Println(model)
	for i := 0; i < rv.NumField(); i++ {
		fmt.Printf("Field[%d]==>%d\n", i, rv.Field(i))
	}
	//model.GetSub("tom")
	//inVal := make([]reflect.Value, 1)
	var inVal []reflect.Value
	inVal = append(inVal, reflect.ValueOf("Tom"))
	rv.MethodByName("GetSub").Call(inVal)

}
