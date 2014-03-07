package main

import (
    "fmt"
    "time"
    "math/rand"
)

func Sleep(amt time.Duration) {
  select {
    case <- time.After(amt):
      return
  }
}

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
        amt := time.Duration(rand.Intn(250))
        Sleep(time.Millisecond * amt)
    }
}


func main() {
    for i := 0; i < 10; i++ {
        go f(i)
    }
    var input string
    fmt.Scanln(&input)
}
