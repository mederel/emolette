package emolette

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var WORDS = map[int][]string{}
var LENGTHS = map[int]int{}

func GenerateLenGoFiles() {
	lines, err := LoadLines()
	CheckErr(err)
	CheckErr(LoadDict(lines, []int{7, 8, 9, 10}))
}

var replaces = map[string]string{
	"é": "e",
	"è": "e",
	"ê": "e",
	"ë": "e",
	"à": "a",
	"á": "a",
	"â": "a",
	"ä": "a",
	"ï": "i",
	"ô": "o",
	"ù": "u",
	"û": "u",
	"æ": "ae",
	"œ": "oe",
}

func LoadDict(lines []string, nbLetters []int) error {
	words := make(map[string]bool)
	wordRegexp, err := regexp.Compile("^([a-zéèêëàáâäïôùûæœ]+)/.*$")
	if err != nil {
		return err
	}
	for _, line := range lines {
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
	log.Print(len(words))
	sortedWords := make([]string, 0, len(words))
	for k, _ := range words {
		sortedWords = append(sortedWords, k)
	}
	sort.Strings(sortedWords)
	LENGTHS[nbLetters] = len(words)
	WORDS[nbLetters] = sortedWords
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
