package prime

import (
  "strconv"
  "strings"
  )

type numeric interface {
  factor() string
}

func newNumber(n number) numeric {
  return &n
}

type number int

func (n *number) factor() string {
  number := int(*n)
  var factors []string

  if number %2 == 0 {
    factors = findFactor(number,2)
    return strings.Join(factors, ",")
  }

  if number %3 == 0 {
    factors = findFactor(number,3)
    return strings.Join(factors, ",")
  }

  return strconv.Itoa(number)
}

func findFactor(number int, lcm int) []string {
  var factor []string
  for IsMoreLCM(number, lcm) && (number%lcm == 0) {
    factor = append(factor, strconv.Itoa(lcm))
    number/=lcm
  }
  factor = append(factor, strconv.Itoa(number))
  return factor
}

func IsMoreLCM(number, lcm int) bool {
  return number/lcm > 1
}
