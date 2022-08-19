package emolette

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var WordsPerLength = map[int][]string{}
var Lengths = []int{7, 8, 9, 10}
var NbWordsPerLength = map[int]int{}

func GenerateLenGoFiles() {
	lines, err := LoadLines()
	CheckErr(err)
	CheckErr(LoadDict(lines))
}

var replaces = map[string]string{
	"é": "e",
	"è": "e",
	"ê": "e",
	"ë": "e",
	"à": "a",
	"â": "a",
	"î": "i",
	"ï": "i",
	"ô": "o",
	"ù": "u",
	"û": "u",
	"ü": "u",
	"æ": "ae",
	"œ": "oe",
	"ç": "c",
}

func LoadDict(lines []string) error {
	wordRegexp, err := regexp.Compile("^([a-zéèêëàâîïôùûüæœç]+)/.*$")
	if err != nil {
		return err
	}
	for _, nbLetters := range Lengths {
		words := make(map[string]bool)
		for _, line := range lines {
			wordRegexp.Match([]byte(strings.ToLower(line)))
			// remove comment lines and uncommon names like Paris or Bill
			if !regexp.MustCompile("^[a-z]").Match([]byte(line)) {
				continue
			}
			// word/<other things> --> keep the word
			word := regexp.MustCompile("^([^/]*)/.*$").ReplaceAllString(line, "$1")
			// remove any none latin character like chinese characters
			if regexp.MustCompile("[^\\p{Latin}]").Find([]byte(word)) != nil {
				continue
			}
			// remove any word containing numbers
			if regexp.MustCompile("\\d").Find([]byte(word)) != nil {
				continue
			}
			// keep only all-lowercase words
			word = strings.ToLower(word)
			// remove all accents
			for k, v := range replaces {
				word = strings.ReplaceAll(word, k, v)
			}
			// keep only words that are 5-letters-big
			if len(word) != nbLetters {
				continue
			}
			words[word] = false
		}
		log.Printf("nb for length %d is %d", nbLetters, len(words))
		sortedWords := make([]string, 0, len(words))
		for k := range words {
			sortedWords = append(sortedWords, k)
		}
		sort.Strings(sortedWords)
		NbWordsPerLength[nbLetters] = len(words)
		WordsPerLength[nbLetters] = sortedWords
	}
	return nil
}

func LoadLines() ([]string, error) {
	content, err := os.ReadFile("fr-toutesvariantes.dic")
	if err != nil {
		return nil, err
	}
	strContent := string(content[:])
	lines := strings.Split(strContent, "\n")
	return lines, nil
}

func CheckErr(err error) {
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}
