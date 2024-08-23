package arrays_and_string_manipulation

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func wordSeek(puzzle [][]string) []string {
	var wordsFound []string
	dictionary := map[string]struct{}{"apple": {}, "cat": {}}

	for _, r := range puzzle {
		for word := range dictionary {
			if strings.Contains(strings.Join(r, ""), word) {
				wordsFound = append(wordsFound, word)
			}
		}
	}

	var sb strings.Builder
	for i := 0; i < 5; i++ {
		for j := range puzzle {
			sb.WriteString(puzzle[j][i])
		}
		for word := range dictionary {
			if strings.Contains(sb.String(), word) {
				wordsFound = append(wordsFound, word)
			}
		}
		sb.Reset()
	}

	return wordsFound
}

func loadFile() ([][]string, error) {
	file, err := os.Open("word_seek.csv")
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)
	return r.ReadAll()
}

func wordSeekFromFile() []string {
	records, err := loadFile()
	if err != nil {
		log.Fatal(err)
	}
	return wordSeek(records)
}
