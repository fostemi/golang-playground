package main

import (
  "fmt"
  "time"
)

func foo(value int) {
  for i := 0; i < value; i++ {
    time.Sleep(time.Second)
    fmt.Println(i)
  }
}

func main () {
  fmt.Println("Starting Goroutines")

  go foo(3)
  go foo(3)

  fmt.Println("Press enter when done")
  fmt.Scanln()
}
