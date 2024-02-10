package maps

import (
	"errors"
)

var ErrNotFound = errors.New("could not find that word")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
