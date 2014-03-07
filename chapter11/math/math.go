package math

// Returns average of a series of numbers
func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total / float64(len(xs))
}

// Returns smallest number in a series
func Min(xs []float64) float64 {
  min := xs[0]
  for _, x := range xs {
    if min > x {
      min = x
    }
  }
  return min
}

// Returns largest number in a series
func Max(xs []float64) float64 {
  max := xs[0]
  for _, x := range xs {
    if max < x {
      max = x
    }
  }
  return max
}
