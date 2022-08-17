package emolette

import (
	"fmt"
	"math/rand"
	"time"
)

func Play() {
	rand.Seed(time.Now().UnixNano())
	indexLength := rand.Intn(len(LENGTHS))
	i := 0
	length := 0
	for k, _ := range LENGTHS {
		if i == indexLength {
			length = k
		}
		i++
	}
	words := WORDS[length]
	theWord := words[rand.Intn(len(words))]
	fmt.Println(theWord)
}
