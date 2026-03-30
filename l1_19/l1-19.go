/*
Разворот строки
Разработать программу, которая переворачивает подаваемую на вход строку.
Например: при вводе строки «главрыба» вывод должен быть «абырвалг».
Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).
*/

package main

import "fmt"

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var s string
	fmt.Scan(&s)

	res := ReverseString(s)
	fmt.Println(res)
}
