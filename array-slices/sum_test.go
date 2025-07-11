package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("slice of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given the array %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// Slices cannot be compared like this: if got != want
	// Careful cause DeepEqual does not care about the types
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
	// From Go 1.21, slices have slices.Equal: if slices.Equal(got, want)
}

func TestSumAllTails(t *testing.T) {

	// Assigning a function to a variable, valid only in this scope
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper() // numero da linha do erro reporta na chamada da função, não aqui dentro
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("all filled slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 4}, []int{0, 9})
		want := []int{6, 9}
		checkSums(t, got, want)
	})

	t.Run("one empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		checkSums(t, got, want)
	})
}
