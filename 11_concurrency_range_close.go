package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

var p interface{}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c {
        fmt.Println(i)
    }
}

// Note: Only the sender should close a channel, never the receiver. Sending on
// a closed channel will cause a panic.

// Another note: Channels aren't like files; you don't usually need to close
// them. Closing is only necessary when the receiver must be told there are no
// more values coming, such as to terminate a range loop.

// you can test if a channel is open or closed too. Closed when ok == false
// v, ok := <-ch
