package main

import (
	"errors"
	"testing"
)


func TestDicionary(t *testing.T) {
	dictionary := Dictionary{}

	dictionary.Add("bomba", "something that explodes")
	
	t.Run("should find added word", func(t *testing.T){
		got,err := dictionary.Search("bomba")
		want := "something that explodes"
		assertError(t, err, nil)
		assertStringEquals(t, got, want)

	})

	t.Run("should not add already existing word", func(t *testing.T){
		err := dictionary.Add("bomba", "something that goes boom")

		assertError(t, err, ErrorAlreadyExists)
	})

	t.Run("failed search not found", func(t *testing.T){
		_,err := dictionary.Search("something")
		want := "could not find the word you were looking for"

		assertError(t, err, ErrorNotFound)

		assertStringEquals(t, err.Error(), want)
	})


}

func assertStringEquals(t testing.TB, got string, want string){
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error){
	t.Helper()

	if(!errors.Is(got, want)){
		t.Errorf("got error %q want %q", got, want)
	}
} 