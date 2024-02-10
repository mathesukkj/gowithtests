package maps

import (
	"mathesukkj/gowithtests/utils"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		utils.AssertCorrectMessageNoParams(t, got, want)

	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("randomWord")

		utils.AssertError(t, got, ErrNotFound)
	})
}
