package main

import (
  "fmt"
)

func fibonacci() func() int {
  old_fib :=-1
  fib := 1
  return func() int {
    new_fib := fib + old_fib
    old_fib = fib
    fib = new_fib
    return fib
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
