package greetings

import (
	// Standard
	"errors"    // Used for creating error instances
	"fmt"       // Used for formatting strings
	"math/rand" // Used for generating random numbers

	// https://pkg.go.dev/
	"rsc.io/quote"
)

// Hello generates a greeting for a named person.
// If the name is empty, it returns an error.
func Hello(name string) (string, error) {
	if name == "" {
		// Return an empty string and an error if no name is provided.
		return "", errors.New("empty name")
	}
	// Generate a message string using a format returned by randomFormat() function.
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// Hellos generates greetings for multiple people.
// It takes a slice of names and returns a map of names to their greetings.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string) // Initialize a map to store name-message pairs.
	for _, name := range names {
		// For each name, retrieve the greeting message.
		message, err := Hello(name)
		if err != nil {
			// If an error occurs, stop processing and return the error.
			return nil, err
		}
		// Store the greeting message in the map using the name as a key.
		messages[name] = message
	}
	return messages, nil
}

// Simple function that calls a dependency
func HelloQuote() {
	fmt.Println(quote.Go())
}

// randomFormat selects a random greeting format.
// It returns a formatted string with a placeholder for the name.
func randomFormat() string {
	// List of possible greeting formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Choose a random index and return the format string at that index.
	return formats[rand.Intn(len(formats))]
}
