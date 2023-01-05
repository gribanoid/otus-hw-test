package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	dict := make(map[string]int)
	arr := strings.Fields(text)
	for _, word := range arr {
		dict[word]++
	}
	words := make([]string, 0, len(dict))
	for w := range dict {
		words = append(words, w)
	}
	sort.Slice(words, func(i, j int) bool {
		if dict[words[i]] == dict[words[j]] {
			return words[j] > words[i]
		}
		return dict[words[i]] > dict[words[j]]
	})
	if len(words) > 10 {
		return words[:10]
	}
	return words
}
