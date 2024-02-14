package main

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

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"wall": "a wall that does wall things"}
	t.Run("new word", func(t *testing.T) {
		dictionary.Add("test", "testing")
		want := "testing"

		got, err := dictionary.Search("test")

		utils.AssertNoError(t, err)
		utils.AssertCorrectMessageNoParams(t, got, want)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "wall"
		err := dictionary.Add("wall", "lmao")

		got, _ := dictionary.Search(word)

		utils.AssertError(t, err, ErrWordExists)
		utils.AssertCorrectMessageNoParams(t, got, "a wall that does wall things")
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"wall": "a wall that does wall things"}
	t.Run("existing word", func(t *testing.T) {
		word := "wall"
		newDefinition := "a brick wall"

		err := dictionary.Update(word, newDefinition)

		got, _ := dictionary.Search(word)

		utils.AssertNoError(t, err)
		utils.AssertCorrectMessageNoParams(t, got, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "lol"
		definition := "lmao"

		err := dictionary.Update(word, definition)

		utils.AssertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"wall": "a wall that does wall things"}
	word := "wall"

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	utils.AssertError(t, err, ErrNotFound)
}
