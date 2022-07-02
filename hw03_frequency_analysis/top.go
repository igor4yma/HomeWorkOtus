package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Spisok struct {
	Word  string
	Count int
}

// var SplitRegexp = regexp.MustCompile(`[\n\t.,!?;:«»()—\"']+`)

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}
	// arr := make([]string, len(strings.Fields(text)))
	// arr = SplitRegexp.Split(text, -1)
	arr := strings.Fields(text)
	cache := make(map[string]int)
	num := 1
	for k := range arr {
		for i := range arr {
			if arr[k] == arr[i] && k != i {
				num++
			}
		}
		_, ok := cache[arr[k]]
		if !ok {
			cache[arr[k]] = num
		}
		num = 1
	}
	result := make([]Spisok, 0)
	for word, cnt := range cache {
		result = append(result, Spisok{Word: word, Count: cnt})
	}
	sort.SliceStable(result, func(i, j int) bool { // sorting lexicographically
		return result[i].Word < result[j].Word
	})
	sort.SliceStable(result, func(i, j int) bool { // sorting by count
		return result[i].Count > result[j].Count
	})
	output := make([]string, 0)
	resultLen := min(10, len(result))
	for _, el := range result[:resultLen] {
		output = append(output, el.Word)
	}
	return output
}

func min(outputLen int, elements ...int) int {
	min := outputLen
	for _, outputLen := range elements {
		if outputLen < min {
			min = outputLen
		}
	}
	return min
}
