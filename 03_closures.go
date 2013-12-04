package main

import "fmt"

// returns a function that returns an int
func fibonacci() func() int {
  old_fib :=-1
  fib := 1

  return func() int {
    fib, old_fib = fib + old_fib, fib
    return fib
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
