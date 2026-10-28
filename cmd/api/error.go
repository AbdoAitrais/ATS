package main



const (
	ErrorNotFound = DictionaryError("could not find the word you were looking for")
 	ErrorAlreadyExists = DictionaryError("word already exists")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}