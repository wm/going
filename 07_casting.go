package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    value := fmt.Sprint(float64(e))
    return "cannot Sqrt negative number: " + value
}

// idiomatic - return (value, nil) or (zero value, error)
func Sqrt(x float64) (float64, error) {
  if x < 0 {
      return 0, ErrNegativeSqrt(x)
  }

  var z = (1.0 + x) / 2.0

  for {
    old_z := z
    z = (z + x/z)/2.0

    if z >= old_z {
      return z, nil
    }
  }

}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
