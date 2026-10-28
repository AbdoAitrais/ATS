package main



type Dictionary map[string]string



func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return  "",ErrorNotFound 
	}
	return definition,nil
}

func (d Dictionary) Add(word string, definition string) error{
	_,err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition 
	case nil:
		return ErrorAlreadyExists
	default:
		return err
	}
	return nil
}