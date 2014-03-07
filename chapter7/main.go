package main

import "fmt"

func average(xs []float64) float64 {
  total := 0.0
  for _, v := range xs {
    total += v
  }
  return total / float64(len(xs))
}

func sum(nums []float64) float64 {
  total := 0.0
  for _, val := range nums {
    total += val
  }
  return total
}

func half(i int) (int, bool) {
  half := i / 2
  return half, half % 2 == 0
}

func max(args ...float64) float64 {
  max := args[0]
  for _, val := range args {
    if val > max {
      max = val
    }
  }
  return max
}

func makeOddGenerator() func() uint {
  i := uint(1)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

func fib(n int) int {
  if n == 0 {
    return 0
  } else if n == 1 {
    return 1
  } else {
    return fib(n-1) + fib(n-2)
  }
}

func main() {
  xs := []float64{98, 93, 77, 82, 83}
  fmt.Println(average(xs))
  fmt.Println(sum(xs[1:3]))
  fmt.Println(half(int(xs[0])))
  fmt.Println(max(xs...))

  nextOdd := makeOddGenerator()
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())

  fmt.Println(fib(10))
}
