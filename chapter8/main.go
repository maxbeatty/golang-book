package main

import "fmt"

func swap(x *int, y *int) {
  tmp := *x
  *x = *y
  *y = tmp
}

func main() {
  x := 1
  y := 2

  swap(&x, &y)

  fmt.Println(x, y)
}
