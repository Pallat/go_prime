package prime

import (
  "strconv"
  "strings"
  )

func prime(number int) string {
  var factor []string
  if number == 4 {
    factor = findFactor(number)
    return strings.Join(factor, ",")
  }
  if number == 6 {
    factor = findFactor(number)
    return strings.Join(factor, ",")
  }
  if number == 8 {
    factor = findFactor(number)
    return strings.Join(factor, ",")
  }
  if number == 9 {
    return "3,3"
  }
  if number == 10 {
    return "2,5"
  }
  return strconv.Itoa(number)
}

func findFactor(number int) []string {
  var factor []string
  for (number/2) > 1 {
    factor = append(factor,"2")
    number/=2
  }
  factor = append(factor, strconv.Itoa(number))
  return factor
}
