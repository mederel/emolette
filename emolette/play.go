package emolette

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Play() {
	rand.Seed(time.Now().UnixNano())
	wordLength := Lengths[rand.Intn(len(Lengths))]
	words := WordsPerLength[wordLength]
	theWord := words[rand.Intn(len(words))]
	log.Printf("chosen word: %s\n", theWord)
	enteredWord := ""
	for enteredWord != theWord {
		log.Printf("rejected solution = %s\n", enteredWord)

		_, err := os.Stdout.WriteString(placeholders.String())
		CheckErr(err)
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			CheckErr(fmt.Errorf("could not understand input"))
		}
		enteredWord = scanner.Text()
	}
	log.Printf("found the soluce!!! %s\n\n", theWord)
}

type CharFeedback int64

const (
	Absent CharFeedback = iota
	Misplaced
	Correct
)

func Feedback(wordLength int, userEntry string, solution string) (string, error) {
	solutionChars := []rune(solution)
	placeholders := strings.Builder{}
	enteredWordChars := []rune(userEntry)
	output := make([]CharFeedback, 0, len(solutionChars))
	for i := 0; i < len(output); i++ {
		output[i] = Absent
	}
	for i := 0; i < wordLength; i++ {
		if output[i] == Absent {
			if enteredWordChars[i] == solutionChars[i] {
				output[i] = Correct
			}
		}
	}
misplaced:
	for i := 0; i < wordLength; i++ {
		if output[i] == Absent {
			for j, solutionChar := range solutionChars {
				if j != i && enteredWordChars[i] == solutionChar {
					output[i] = Misplaced
					break misplaced
				}
			}
		}
	}
	placeholders.WriteString("\n")
}
