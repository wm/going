package main

import (
    "code.google.com/p/go-tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
    walk(t, ch)
    close(ch)
}

func walk(t *tree.Tree, ch chan int) {
  if t != nil {
    walk(t.Right, ch)
    ch <- t.Value
    walk(t.Left, ch)
  }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1, ch2 := make(chan int), make(chan int)

  go Walk(t1, ch1)
  go Walk(t2, ch2)

  for v1 := range ch1 {
    v2 := <-ch2
    fmt.Println([]int{v2, v1})
    if v1 != v2 {
      return false
    }
  }
  return true
}

func main() {
  fmt.Println(Same(tree.New(1), tree.New(1)))
  fmt.Println(Same(tree.New(3), tree.New(1)))
}
