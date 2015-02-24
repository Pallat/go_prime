package prime

import (
  "strconv"
  "strings"
  )

type number int

func (n *number) factor() string {
  return prime(int(*n))
}

func prime(number int) string {
  var factor []string
  if number %2 == 0 {
    factor = findFactor(number,2)
    return strings.Join(factor, ",")
  }

  if number %3 == 0 {
    factor = findFactor(number,3)
    return strings.Join(factor, ",")
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
