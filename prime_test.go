package prime

import (
  "testing"
  )

func TestPrimeOneShouldBeOne(t *testing.T) {
  factor := prime(1)
  if factor != "1" {
    t.Error("It should be 1 but got ", factor)
  }
}
