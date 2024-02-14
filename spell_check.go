package spell_check

import (
	"bufio"
	"os"
	"time"
)

func init() {
	start_time := time.Now()
	file, err := os.Open(DictionaryPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		dictionary.AddWord(word)
	}

	time_taken := time.Since(start_time)

	println("Time taken to load dictionary: ", time_taken.Seconds())
	println("Size of dictionary: ", len(dictionary.words))
}

func SpellCheck(inputFilePath, outputFilePath string) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		word := scanner.Text()
		correctedWord := dictionary.CheckWord(word)
		writer.WriteString(correctedWord + "\n")
	}

	writer.Flush()
}

// map reduce? 😳
func SpellCheckMultithreaded(inputFilePath, outputFilePath string) {
}
