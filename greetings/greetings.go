package greetings

import (
  "fmt"
  "errors"
  "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
  // If no name was given, return an error with a message.
  if name == "" {
    return "", errors.New("empty name")
  }

  // Return a greeting that embeds the name in a message.
  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}

// randomFormat returns one of a set of greetings. 
// The returned message is selected randomly.
func randomFormat() string {
  // A slice of message formats.
  formats := []string{
    "Hi, %v. Welcome!",
    "Great to see you, %v.",
    "What's up, %v?",
    "Howdy, %v. How goes it?",
    "Well met, %v.",
  }

  // Select a random index.
  return formats[rand.Intn(len(formats))]
}
