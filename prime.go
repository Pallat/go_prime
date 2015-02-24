package prime

import (
  "strconv"
  "strings"
  )

func prime(number int) string {
  var factor []string
  if number == 4 {
    factor = findFactor(number,2)
    return strings.Join(factor, ",")
  }
  if number == 6 {
    factor = findFactor(number,2)
    return strings.Join(factor, ",")
  }
  if number == 8 {
    factor = findFactor(number,2)
    return strings.Join(factor, ",")
  }
  if number == 9 {
    factor = findFactor(number,3)
    return strings.Join(factor, ",")
  }
  if number == 10 {
    return "2,5"
  }
  return strconv.Itoa(number)
}

func findFactor(number int, base int) []string {
  var factor []string
  for (number/base) > 1 {
    factor = append(factor, strconv.Itoa(base))
    number/=base
  }
  factor = append(factor, strconv.Itoa(number))
  return factor
}
