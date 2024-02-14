package spell_check

import (
	"testing"
)

func TestCheckWord(t *testing.T) {
	dictionary.AddWord("hello")

	correctedWord := dictionary.CheckWord("hellw")

	if correctedWord != "hello" {
		t.Errorf("Expected hello, got %s", correctedWord)
	}

	correctedWord = dictionary.CheckWord("helo")
	if correctedWord != "hello" {
		t.Errorf("Expected hello, got %s", correctedWord)
	}

	correctedWord = dictionary.CheckWord("hell")
	if correctedWord != "hell" {
		t.Errorf("Expected hello, got %s", correctedWord)
	}
}

func TestGenerateData(t *testing.T) {
	GenerateRandomData("raw/list1k.txt", 1000)
}
