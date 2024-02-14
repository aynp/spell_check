package spell_check

import (
	"testing"
)

var raw1k = "raw/list1k.txt"
var raw10k = "raw/list10k.txt"
var raw100k = "raw/list100k.txt"

var output1k = "output/list1k.txt"
var output10k = "output/list10k.txt"
var output100k = "output/list100k.txt"

func init() {
	GenerateRandomData(raw1k, 1000)
	GenerateRandomData(raw10k, 10000)
	GenerateRandomData(raw100k, 100000)
}

func BenchmarkSpellCheck1k(b *testing.B) {
	SpellCheck(raw1k, output1k)
}

func BenchmarkSpellCheck10k(b *testing.B) {
	SpellCheck(raw10k, output10k)
}

func BenchmarkSpellCheckMultithreaded10k(b *testing.B) {
	SpellCheckMultithreaded(raw10k, output10k)
}

func BenchmarkSpellCheckMultithreaded100k(b *testing.B) {
	SpellCheckMultithreaded(raw100k, output100k)
}
