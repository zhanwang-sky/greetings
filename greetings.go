package greetings

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

var formats = []string{
	"Hi, %v. Welcome!",
	"Great to see you, %v!",
	"Hail, %v! Well met!",
}

func randomFormat() string {
	return formats[rand.IntN(len(formats))]
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
