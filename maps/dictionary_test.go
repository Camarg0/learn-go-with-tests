package main

import (
	"testing"
)

const (
	word       = "test"
	definition = "this is just a test"
)

func TestDictionary(t *testing.T) {
	t.Run("word found", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		got, _ := dictionary.Search(word)
		want := definition

		assertStrings(t, got, want)
	})

	t.Run("word not found", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		_, err := dictionary.Search("notfound")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding an item", func(t *testing.T) {
		key := word
		value := definition
		dictionary := Dictionary{}

		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("trying to add a duplicate item", func(t *testing.T) {
		key := word
		value := definition
		dictionary := Dictionary{key: value}

		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrorAlreadyExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := word
		value := definition
		dictionary := Dictionary{key: value}
		update := "test updated"

		dictionary.Update(key, update)

		assertDefinition(t, dictionary, key, update)
	})

	t.Run("non existing word", func(t *testing.T) {
		key := word
		value := definition
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t, err, ErrorDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := word
		value := definition
		dictionary := Dictionary{key: value}

		err := dictionary.Delete(key)

		assertError(t, err, nil)

		_, err = dictionary.Search(key)

		assertError(t, err, ErrorNotFound)
	})

	t.Run("non existing word", func(t *testing.T) {
		key := word
		dictionary := Dictionary{}

		err := dictionary.Delete(key)

		assertError(t, err, ErrorDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal(err)
	}
	assertStrings(t, got, value)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
