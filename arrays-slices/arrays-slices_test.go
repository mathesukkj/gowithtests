package main

import (
	"fmt"
	"mathesukkj/gowithtests/utils"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("arrays with 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 60, 4}
		got := Sum(numbers)
		want := 70

		utils.AssertCorrectMessageWithParam(t, got, want, numbers)
	})
	t.Run("collections of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		utils.AssertCorrectMessageWithParam(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 4}, []int{2, 3, 5})
	want := []int{5, 10}

	utils.AssertCorrectMessageSlice(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 45}, []int{3, 6, 1}, []int{1})
		want := []int{45, 1, 1}

		utils.AssertCorrectMessageSlice(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 6})
		want := []int{0, 6}

		utils.AssertCorrectMessageSlice(t, got, want)
	})
}

func ExampleSum() {
	numbers := []int{1, 6, 20, 0, 0}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 27
}
