package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("Could not find the word you're looking...")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
