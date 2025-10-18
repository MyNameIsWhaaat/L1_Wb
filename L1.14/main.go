package main

import (
	"fmt"
	"reflect"
)

func detectType(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan any:
		return "chan"
	default:
		if reflect.ValueOf(x).IsValid() && reflect.ValueOf(x).Kind() == reflect.Chan {
			return "chan"
		}
		return "unknown"
	}
}

func main() {
	fmt.Println(detectType(42))
	fmt.Println(detectType("hi"))
	fmt.Println(detectType(true))

	chInt := make(chan int)
	fmt.Println(detectType(chInt))

	chStruct := make(chan struct{})
	fmt.Println(detectType(chStruct))

	var recvOnly <-chan int = chInt
	fmt.Println(detectType(recvOnly))
}
