package iteration

import (
	"errors"
	"strings"
)

var ErrNegativeNumbers = errors.New("Negative Numbers are not allowed")

func Repeat(character string, counter int) (string, error) {
	if counter < 0 {
		return "", ErrNegativeNumbers
	}

	return strings.Repeat(character, counter), nil
}
