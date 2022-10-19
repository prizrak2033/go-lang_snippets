package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  // seed to get different result every time
  rand.Seed(time.Now().UnixNano())
  fmt.Println(rand.Intn(50))
  fmt.Println(rand.Float64())
}
