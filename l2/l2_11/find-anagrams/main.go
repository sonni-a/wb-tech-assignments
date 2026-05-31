package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	groups := make(map[string][]string)
	firstWord := make(map[string]string)

	for _, word := range words {
		lower := strings.ToLower(word)

		runes := []rune(lower)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		key := string(runes)

		groups[key] = append(groups[key], lower)

		if _, exists := firstWord[key]; !exists {
			firstWord[key] = lower
		}
	}

	result := make(map[string][]string)

	for key, group := range groups {
		if len(group) < 2 {
			continue
		}

		sort.Strings(group)
		result[firstWord[key]] = group
	}

	return result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	fmt.Println(findAnagrams(words))
}
