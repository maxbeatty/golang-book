package main

import "fmt"

func main() {
  fmt.Print("Enter Fehrenheit: ")
  var input float64
  fmt.Scanf("%f", &input)

  Celcius := (input - 32) * 5/9

  fmt.Println(Celcius)
}
