package main

import (
  "fmt"
  "sync"
  "time"
)

func foo(wg * sync.WaitGroup) {
  time.Sleep(time.Second)
  fmt.Println("Finished foo")

  wg.Done()
}

func main() {
  fmt.Println("Waitgroups tutorial")
  
  var wg sync.WaitGroup
  
  wg.Add(1)
  go foo(&wg)
  wg.Wait()

  fmt.Println("Finished main")
}
