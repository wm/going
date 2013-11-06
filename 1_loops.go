package main

import (
  "fmt"
)

func Sqrt(x float64) float64 {
  z := (1.0 + x) / 2.0

  fmt.Println("Type 4")
  for {
    old_z := z
    z = (z + x/z)/2.0
    if z >= old_z {
      return z
    }
  }
}

func main() {
  var pow = []int{1, 2, 4, 8, 16}
  // pow := []int{1, 2, 4, 8, 16}

  fmt.Println("Type 1")
  for i, value := range pow {
    fmt.Println(i, value)
  }

  fmt.Println("Type 2")
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }

  fmt.Println("Type 3")
  i := 0 // declare because previous was not in scope
  for i < 5 {
    i++
    fmt.Println(i)
  }

  fmt.Println(Sqrt(256))
}
