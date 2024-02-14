package spell_check

type Dictionary struct {
	words map[string]string
}

var dictionary = Dictionary{
	words: make(map[string]string),
}

func (d *Dictionary) AddWord(word string) {
	dictionary.words[word] = word

  // single extra character
	for i := 0; i <= len(word); i++ {
			temp_word := word[:i] + "#" + word[i:]
			dictionary.words[temp_word] = word
	}

  // single chracter wrong
	for i := 0; i < len(word); i++ {
		temp_word := word[:i] + "#" + word[i+1:]
		dictionary.words[temp_word] = word
	}

  // [duplicated] included below
	// two adjacent characters swapped
	// for i := 0; i < len(word)-1; i++ {
	// 	temp_word := word[:i] + "##" + word[i+2:]
	// 	println(temp_word)
	// 	dictionary.words[temp_word] = word
	// }

  // two characters wrong
  for i := 0; i <= len(word)-1; i++ {
    for j := i+1; j < len(word); j++ {
      temp_word := word[:i] + "#" + word[i+1:j] + "#" + word[j+1:]
      dictionary.words[temp_word] = word
    }
  }
}

func (d *Dictionary) CheckWord(word string) string {
	if _, ok := dictionary.words[word]; ok {
		return word
	}

	// single extra character
	for i := 0; i <= len(word); i++ {
		temp_word := word[:i] + "#" + word[i:]
		if _, ok := dictionary.words[temp_word]; ok {
			return dictionary.words[temp_word]
		}
	}

	// single chracter wrong
	for i := 0; i < len(word); i++ {
		temp_word := word[:i] + "#" + word[i+1:]
		if _, ok := dictionary.words[temp_word]; ok {
			return dictionary.words[temp_word]
		}
	}

	// two characters wrong
	for i := 0; i < len(word)-1; i++ {
		for j := i + 1; j < len(word); j++ {
			temp_word := word[:i] + "#" + word[i+1:j] + "#" + word[j+1:]
			if _, ok := dictionary.words[temp_word]; ok {
				return dictionary.words[temp_word]
			}
		}
	}

	return "~~~"
}
