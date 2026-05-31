/*
Уникальные символы в строке
Разработать программу, которая проверяет, что все символы в строке встречаются один раз
(т.е. строка состоит из уникальных символов).
Вывод: true, если все символы уникальны, false, если есть повторения.
Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.
Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.
Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isAllUnique(s string) bool {
	s = strings.ToLower(s)

	seen := make(map[rune]struct{})

	for _, ch := range s {
		if _, ok := seen[ch]; ok {
			return false
		}
		seen[ch] = struct{}{}
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	fmt.Println(isAllUnique(s))
}
