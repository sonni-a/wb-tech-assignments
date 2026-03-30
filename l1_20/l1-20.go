/*
Разворот слов в предложении
Разработать программу, которая переворачивает порядок слов в строке.
Пример: входная строка:
«snow dog sun», выход: «sun dog snow».
Считайте, что слова разделяются одиночным пробелом.
Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».
*/
package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	result := ReverseWords(input)
	fmt.Println(result)
}
