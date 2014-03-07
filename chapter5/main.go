package main

import "fmt"

func main() {
/*  i := 1

  for i<=10 {
    fmt.Println(i)
    i += 1
  }

  for i = i; i <= 20; i++ {
    fmt.Println(i)
  }

  for j := 1; j <= 10; j++ {
    if j % 2 == 0 {
      fmt.Println(j, "even")
    } else {
      fmt.Println(j, "odd")
    }
  }*/

  for i := 1; i < 101; i++ {
    var x string = ""

    if i % 3 == 0 {
      x = "Fizz"
    }

    if i % 5 == 0 {
      x += "Buzz"
    }

    if len(x) > 0 {
      fmt.Println(x)
    }
  }
}
