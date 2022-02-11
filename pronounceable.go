package pronounceable

import (
	"math"
	"os"
	"strings"
)

type Dataset map[int]map[string]float64

func NewDatasetFromWords(words []string) Dataset {
	dataset := make(Dataset)

	for _, word := range words {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}

		for i := 0; i <= 3; i++ {
			ngrams := combinations(word, i)

			for _, ngram := range ngrams {
				if _, ok := dataset[i]; !ok {
					dataset[i] = make(map[string]float64)
				}

				dataset[i][ngram]++
			}
		}
	}

	return dataset
}

func NewDatasetFromFile(source string) (Dataset, error) {
	bytes, err := os.ReadFile(source)
	if err != nil {
		return nil, err
	}

	words := strings.Split(string(bytes), "\n")

	return NewDatasetFromWords(words), nil
}

func (d Dataset) Score(sequence string) float64 {
	var score float64
	sequence = strings.ToLower(sequence)

	for i := 0; i <= 3; i++ {
		if score == math.Inf(-1) {
			break
		}

		ngrams := combinations(sequence, i)

		for _, ngram := range ngrams {
			count := d[i][ngram] + 1

			score += math.Log(count/float64(len(d[i]))) * float64(5+i)
		}
	}

	score /= 1.5 * float64(len(sequence))

	if score > 100 {
		return 1
	}

	return score / 100
}

// combinations returns all possible ngrams of a given word.
func combinations(word string, n int) []string {
	if strings.Contains(word, " ") {
		panic("word cannot contain spaces")
	}

	var ngrams []string
	for i := 0; i < len(word)-n+1; i++ {
		if i+n <= len(word) {
			ngrams = append(ngrams, word[i:i+n])
		}
	}

	return ngrams
}
