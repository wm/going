// http://blog.golang.org/json-and-go
// Make some examples from the above
// try and incorporate a select with channels and goroutines
package main

import (
    "encoding/json"
    "log"
    "os"
)

func main() {
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }
        for k := range v {
            if k != "Name" {
                delete(v, k)
            }
        }
        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
    }
}
