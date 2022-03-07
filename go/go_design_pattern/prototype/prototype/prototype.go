package prototype

import (
	"encoding/json"
)

type Keyword struct {
	Word  string
	Visit int
	// UpdatedAt *time.Time
}

func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	// json Unmarshal time.Time 会损失精度
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

type Keywords map[string]*Keyword

func (words Keywords) Clone(updatedWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range words {
		newKeywords[k] = v
	}

	for _, word := range updatedWords {
		newKeywords[word.Word] = word.Clone()
	}

	return newKeywords
}
