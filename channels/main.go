package main

import (
  "fmt"
  "time"
)

func SendValue(c chan string) {
  time.Sleep(1 * time.Second)
  c <- "hello"
}

func main() {
  fmt.Println("Channels tutorial")

  values := make(chan string)

  defer close(values)
  
  go SendValue(values)
  go SendValue(values)
  
  value := <-values
  fmt.Println(value)

  time.Sleep(1 * time.Second)
}
