package math

import "testing"

type testpair struct {
  values []float64
  expected float64
}

var avgTests = []testpair{
  { []float64{1,2}, 1.5 },
  { []float64{1,1,1,1,1}, 1 },
  { []float64{-1,1}, 0 },
}

func TestAverage(t *testing.T) {
  for _, pair := range avgTests {
    v := Average(pair.values)
    if v != pair.expected {
      t.Error(
        "For", pair.values,
        "expected", pair.expected,
        "got", v,
      )
    }
  }
}

var minTests = []testpair{
  { []float64{1,2}, 1 },
  { []float64{2,1}, 1 },
  { []float64{3,1,2}, 1 },
  { []float64{-1,-2,-3}, -3},
}

func TestMin(t *testing.T) {
  for _, pair := range minTests {
    v := Min(pair.values)
    if v != pair.expected {
      t.Error(
        "For", pair.values,
        "expected", pair.expected,
        "got", v,
      )
    }
  }
}
