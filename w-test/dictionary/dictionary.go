package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find")

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}
