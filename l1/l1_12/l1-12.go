/*
Собственное множество строк
Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.
Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.
*/
package main

import "fmt"

func uniqueWords(words []string) []string {
	unique := make(map[string]struct{})
	result := []string{}

	for _, word := range words {
		if _, ok := unique[word]; !ok {
			unique[word] = struct{}{}
			result = append(result, word)
		}
	}

	return result
}

func main() {
	fmt.Println(uniqueWords([]string{"cat", "cat", "dog", "cat", "tree"}))
}
