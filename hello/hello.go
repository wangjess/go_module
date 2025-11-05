package main

import (
  "fmt"
  "log"

  "example.com/greetings"
)

func main() {
  // Set properties of the predefined logger, including
  // the log entry prefix and a flag to disable printing
  // the time, source file, and line number.
  log.SetPrefix("greetings: ")
  log.SetFlags(0)

  // A slice of names.
  names := []string{"Gladys", "Samantha", "Darrin"}

  // Get a greeting message and print it.
  messages, err := greetings.Hellos(names)
  
  // If an error was returned, print it to the console.
  // Exit the program.
  if err != nil {
    log.Fatal(err)
  }

  // If no error, print the returned message to console.
  fmt.Println(messages)
}
