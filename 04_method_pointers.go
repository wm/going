package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

// mutator - idiomatic
func (v *Vertex) MScale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func (v Vertex) Scale(f float64) Vertex {
  v.X = v.X * f
  v.Y = v.Y * f
  return v
}

// no mutation. If not pointer then a copy is passed in
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MDouble(x *int) int {
  *x = 2*(*x)
  return *x
}

func Double(x int) int {
  x = 2*x
  return x
}

func main() {
  fmt.Println("Double(7) =>", Double(7))

  x := 7
  Double(x)
  fmt.Println("x := 7; Double(x);  x => ", x)

  y := 7
  fmt.Println("y := 7; MDouble(y); y => ", MDouble(&y))
  fmt.Println()

  // methods on objects?
  v := Vertex{3, 4}
  fmt.Println("             v = ", v, "  ;        v.Abs() => ", v.Abs())

  v.Scale(5)
  fmt.Println("v.Scale(5);  v = ", v, "  ;        v.Abs() => ", v.Abs())

  scaled_v := v.Scale(5)
  fmt.Println("      scaled_v = ", scaled_v, "; scaled_v.Abs() => ", scaled_v.Abs())

  v.MScale(5)
  fmt.Println("v.MScale(5); v = ", v, ";        v.Abs() => ", v.Abs())
  // see http://stackoverflow.com/questions/15096329/golang-pointers for
  // explanation of pointer/value methods and their generated/looked up
  // counterparts
}
