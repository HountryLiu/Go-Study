package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `bili:"bili_name",cili:"cili_name"`
	Sex  string `bili:"bili_sex",cili:"cili_sex"`
}

func printTag(ptr interface{}) {
	reType := reflect.TypeOf(ptr)
	if reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct {
		panic("param is not the pointer of struct!")
	}
	reVal := reflect.ValueOf(ptr).Elem()

	for i := 0; i < reVal.NumField(); i++ {
		field := reVal.Type().Field(i)
		fmt.Println(field.Tag.Get("bili"))
	}
}
func main() {
	user := User{
		Name: "n",
		Sex:  "s",
	}

	printTag(&user)
}
