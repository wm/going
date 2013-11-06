package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertex) Scale(f float64) Vertex {
  v.X = v.X * f
  v.Y = v.Y * f
  return v
}

// mutator
func (v *Vertex) MScale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  v.Scale(5)
  fmt.Println(v, v.Abs())

  other_v := v.Scale(5)
  fmt.Println(other_v, other_v.Abs())

  v.MScale(5)
  fmt.Println(v, v.Abs())
}
