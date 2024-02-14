package spell_check

import (
	"math/rand"
	"os"
)

// GenerateRandomData generates random words and writes them to a file. Generates absolutely gibberish words
// but that's okay for as spell check would be closer to the worst case scenario.
func GenerateRandomData(outputFilePath string, numWords int) {
	file, err := os.Create(outputFilePath)

	if err != nil {
		panic(err)
	}

	for i := 0; i < numWords; i++ {
		word := ""
		for j := 0; j < 2 + rand.Intn(16); j++ {
			word += string(rune('a' + rand.Intn(26)))
		}
		file.WriteString(word + "\n")
	}
}
