package main

import (
	//"encoding/json"
	"fmt"
	//"reflect"
	//"time"
)

func main() {
	a := map[string]interface{}{"S": 1, "H": "A"}
	fmt.Printf("%+v\n", a)
	//fmt.Printf("\033[96m  \r%+v\n", now.Unix())
	//fmt.Printf("\033[96m  \r%+v\n", reflect.TypeOf(now.Unix()))
}
