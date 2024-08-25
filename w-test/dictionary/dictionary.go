package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound      = errors.New("could not find")
	ErrAlreadyExists = errors.New("already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}

	return nil
}
