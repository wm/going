package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var medical map[string]Vertex

var computer = map[string]Vertex{
  "Bell Labs": Vertex{ 40.68433, -74.39967 },
  "Google": Vertex{ 37.42202, -122.08408 },
}

// var computer = map[string]Vertex{
//   "Bell Labs": {40.68433, -74.39967},
//   "Google":    {37.42202, -122.08408},
// }

func main() {
  medical = make(map[string]Vertex)
  medical["Iora Health"] = Vertex{
    40.68433, -74.39967,
  }

  fmt.Println(medical["Iora Health"])
  fmt.Println(computer["Bell Labs"])

  delete(medical, "Iora Health")
  fmt.Println("The value:", medical["Iora Health"])

  v, ok := medical["Iora Health"]
  fmt.Println("The value:", v, "Present?", ok)

  v, ok = computer["Bell Labs"]
  fmt.Println("The value:", v, "Present?", ok)
}
