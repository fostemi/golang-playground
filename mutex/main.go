package main

import (
  "fmt"
  "sync"
)

var (
  mutex sync.Mutex
  balance int
)

func deposit( value int, wg * sync.WaitGroup ) {
  
  mutex.Lock()
  fmt.Printf("Deposit %d to account with balance %d\n", value, balance)
  balance += value
  mutex.Unlock()
  
  wg.Done()

}

func withdrawal( value int, wg * sync.WaitGroup ) {
  
  mutex.Lock()
  fmt.Printf("Withdrawal %d to account with balance %d\n", value, balance)
  balance -= value
  mutex.Unlock()

  wg.Done()

}

func main() {
  fmt.Println("Starting mutex 101")
  
  balance = 1000
  
  var wg sync.WaitGroup
  wg.Add(2)
  
  go deposit(700, &wg)
  go withdrawal(500, &wg)

  wg.Wait()

  fmt.Println("Balance: %d", balance)

}
