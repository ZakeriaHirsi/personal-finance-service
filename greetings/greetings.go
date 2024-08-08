package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello greeting returned for a named person

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	format := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v!",
		"Hail, %v! Well met!",
	}

	return format[rand.Intn(len(format))]
}
