package main

import (
	"fmt"
	"mathesukkj/gowithtests/utils"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Matheus", "")
		want := "Hello, Matheus!"
		utils.AssertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!', when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		utils.AssertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello' but in Portuguese", func(t *testing.T) {
		got := Hello("Matheus", "PT-BR")
		want := "Oieee, Matheus!"
		utils.AssertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello' but in french", func(t *testing.T) {
		got := Hello("Matheus", "FR")
		want := "Bonjour, Matheus!"
		utils.AssertCorrectMessage(t, got, want)
	})
}

func ExampleHello() {
	str := Hello("Matheus", "PT-BR")
	fmt.Println(str)
	// Output: Oieee, Matheus!
}
