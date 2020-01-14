package main

import (
	//"encoding/json"
	"fmt"
	"reflect"
)

type Animal struct {
	Name  string
	Order string
}

//func (a *Animal) fill(data map[string]interface{}) {
//	t := reflect.ValueOf(result).Elem()
//}

func main() {
	ex := map[string]int{"a": 1, "b": 2}
	j, k := ex["b"]
	fmt.Println(j, k)
	a := &Animal{}
	structValue := reflect.ValueOf(a).Elem()
	//structFieldValue := structValue.FieldByName("Name")
	fmt.Printf("%+v\n", structValue)
	fmt.Printf("%+v\n", reflect.TypeOf(structValue))
}
