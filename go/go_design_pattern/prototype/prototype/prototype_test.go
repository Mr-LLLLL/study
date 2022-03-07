package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyword_Clone(t *testing.T) {
	// updateAt, _ := time.Parse("2006", "2020")
	words := Keywords{
		"testA": &Keyword{
			Word:  "testA",
			Visit: 1,
			// UpdatedAt: &updateAt,
		},
		"testB": &Keyword{
			Word:  "testB",
			Visit: 2,
			// UpdatedAt: &updateAt,
		},
		"testC": &Keyword{
			Word:  "testC",
			Visit: 3,
			// UpdatedAt: &updateAt,
		},
	}

	// now := time.Now()
	updatedWords := []*Keyword{
		{
			Word:  "testB",
			Visit: 3,
			// UpdatedAt: &now,
		},
	}

	got := words.Clone(updatedWords)

	assert.Equal(t, words["testA"], got["testA"])
	assert.NotEqual(t, words["testB"], got["testB"])
	assert.Equal(t, updatedWords[0], got["testB"])
	assert.Equal(t, words["testC"], got["testC"])
}
