package prime

import (
  "testing"
  )

func TestPrimeOneShouldBeOne(t *testing.T) {
  var prime number = 1
  factor := prime.factor()
  if factor != "1" {
    t.Error("It should be 1 but got ", factor)
  }
}

func TestPrimeTwoShouldBeTwo(t *testing.T) {
  var prime number = 2
  factor := prime.factor()
  if factor != "2" {
    t.Error("It should be 2 but got ", factor)
  }
}

func TestPrimeThreeShouldBeThree(t *testing.T) {
  var prime number = 3
  factor := prime.factor()
  if factor != "3" {
    t.Error("It should be 3 but got ", factor)
  }
}

func TestPrimeFourShouldBeTwoTwo(t *testing.T) {
  var prime number = 4
  factor := prime.factor()
  if factor != "2,2" {
    t.Error("It should be 2,2 but got ", factor)
  }
}

func TestPrimeFiveShouldBeFive(t *testing.T) {
  var prime number = 5
  factor := prime.factor()
  if factor != "5" {
    t.Error("It should be 5 but got ", factor)
  }
}

func TestPrimeSixShouldBeTwoThree(t *testing.T) {
  var prime number = 6
  factor := prime.factor()
  if factor != "2,3" {
    t.Error("It should be 2,3 but got ", factor)
  }
}

func TestPrimeEightShouldBeTwoTwoTwo(t *testing.T) {
  var prime number = 8
  factor := prime.factor()
  if factor != "2,2,2" {
    t.Error("It should be 2,2,2 but got ", factor)
  }
}

func TestPrimeNineShouldBeThreeThree(t *testing.T) {
  var prime number = 9
  factor := prime.factor()
  if factor != "3,3" {
    t.Error("It should be 3,3 but got ", factor)
  }
}

func TestPrimeTenShouldBeTwoFive(t *testing.T) {
  var prime number = 10
  factor := prime.factor()
  if factor != "2,5" {
    t.Error("It should be 2,5 but got ", factor)
  }
}

func TestPrimeTwelveShouldBeTwoTwoThree(t *testing.T) {
  var prime number = 12
  factor := prime.factor()
  if factor != "2,2,3" {
    t.Error("It should be 2,2,3 but got ", factor)
  }
}

func TestPrimeForteenShouldBeTwoSeven(t *testing.T) {
  var prime number = 14
  factor := prime.factor()
  if factor != "2,7" {
    t.Error("It should be 2,7 but got ", factor)
  }
}

func TestPrimeFifteenShouldBeThreeFive(t *testing.T) {
  var prime number = 15
  factor := prime.factor()
  if factor != "3,5" {
    t.Error("It should be 3,5 but got ", factor)
  }
}
