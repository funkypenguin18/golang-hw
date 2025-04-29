package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

// WordFrequency хранит слово и частоту его появления.
type WordFrequency struct {
	Word  string
	Count int
}

func Top10(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	// сначала найдём все слова, разделив текст по пробельным символам
	words := strings.Fields(text)

	// подсчитываем частоту слов
	freqMap := make(map[string]int)
	for _, word := range words {
		freqMap[word]++
	}

	// преобразуем map в слайс структур для сортировки
	frequencies := make([]WordFrequency, 0, len(freqMap))
	for word, count := range freqMap {
		frequencies = append(frequencies, WordFrequency{Word: word, Count: count})
	}

	// сортируем сначала по частоте (по убыванию), затем лексикографически (по возрастанию)
	sort.Slice(frequencies, func(i, j int) bool {
		if frequencies[i].Count == frequencies[j].Count {
			return frequencies[i].Word < frequencies[j].Word
		}
		return frequencies[i].Count > frequencies[j].Count
	})

	// выбираем первые 10 слов
	top := make([]string, 0, 10)
	for i := 0; i < len(frequencies) && i < 10; i++ {
		top = append(top, frequencies[i].Word)
	}

	return top
}
