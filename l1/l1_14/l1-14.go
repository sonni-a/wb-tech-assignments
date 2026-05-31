/*
Определение типа переменной в runtime
Разработать программу, которая в runtime способна определить тип переменной,
переданной в неё (на вход подаётся interface{}).
Типы, которые нужно распознавать: int, string, bool, chan (канал).
Подсказка: оператор типа switch v.(type) поможет в решении.
*/
package main

import (
	"fmt"
	"reflect"
)

func typeRuntime(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			return "chan"
		}
		return "unknown"
	}
}

func main() {
	ch := make(chan int)
	fmt.Println(typeRuntime(ch))
}
