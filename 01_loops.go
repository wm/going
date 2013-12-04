// You should name the main package 'main'
package main

// how to import other packages (self defined or standard lib)
import (
  "fmt"
)

// defining a function named Sqrt
func Sqrt(x float64) float64 {
  var z = (1.0 + x) / 2.0

  fmt.Println("endless looping")
  for {
    old_z := z
    z = (z + x/z)/2.0
    if z >= old_z {
      return z
    }
  }
}

func main() {

  // var pow = make([]int, 5)
  // var pow = []int{1, 2, 4, 8, 16}
  pow := []int{1, 2, 4, 8, 16}

  fmt.Println("basic for loop (like C/Java) all variables need to be used")
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }

  fmt.Println("loop with a range returns index, value")
  for i, value := range pow {
    fmt.Println(i, value)
  }

  i := 0 // declare because previous was not in scope

  fmt.Println("While loop")
  for i < 5 {
    i++
    fmt.Println(i)
  }

  fmt.Println(Sqrt(256))
}
