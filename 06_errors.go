package main

import (
  "fmt"
  "time"
)

// already defined in go
type error interface {
  Error() string
}

type MyError struct {
  When time.Time
  What string
}

func (e *MyError) Error() string {
  return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
  return &MyError{
    time.Now(),
    "it didn't work",
  }
}

func main() {
  if err := run(); err != nil {
    fmt.Println(err)
  }
}

// Interfaces are satisified implicitly
// we return a pointer to MyError but we could return a MyError
// we return a pointer to an error but we could return an error
