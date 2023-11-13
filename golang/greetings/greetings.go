package greetings
import (
	"fmt"
	"errors"
	"math/rand"
)

// nil is the zero value for pointers, interfaces, maps, slices, channels and function types, representing an uninitialized value
// := is a short variable declaration

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string] string, error) {
	messages := make(map[string] string)
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
	// a slice of message formats
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message format by specifying a random index for the slice
	return formats[rand.Intn(len(formats))]
}