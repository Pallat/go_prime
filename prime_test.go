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

func TestPrimeTwoShouldBeTwo(t *testing.T) {
  factor := prime(2)
  if factor != "2" {
    t.Error("It should be 2 but got ", factor)
  }
}

func TestPrimeThreeShouldBeThree(t *testing.T) {
  factor := prime(3)
  if factor != "3" {
    t.Error("It should be 3 but got ", factor)
  }
}

func TestPrimeFourShouldBeTwoTwo(t *testing.T) {
  factor := prime(4)
  if factor != "2,2" {
    t.Error("It should be 2,2 but got ", factor)
  }
}

func TestPrimeFiveShouldBeFive(t *testing.T) {
  factor := prime(5)
  if factor != "5" {
    t.Error("It should be 5 but got ", factor)
  }
}

func TestPrimeSixShouldBeTwoThree(t *testing.T) {
  factor := prime(6)
  if factor != "2,3" {
    t.Error("It should be 2,3 but got ", factor)
  }
}

func TestPrimeEightShouldBeTwoTwoTwo(t *testing.T) {
  factor := prime(8)
  if factor != "2,2,2" {
    t.Error("It should be 2,2,2 but got ", factor)
  }
}

func TestPrimeNineShouldBeThreeThree(t *testing.T) {
  factor := prime(9)
  if factor != "3,3" {
    t.Error("It should be 3,3 but got ", factor)
  }
}
