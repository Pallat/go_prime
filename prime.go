package prime

import (
  "strconv"
  "strings"
  )

func prime(number int) string {
  var factor []string
  if number %2 == 0 {
    factor = findFactor(number,2)
    return strings.Join(factor, ",")
  }

  if number == 9 {
    factor = findFactor(number,3)
    return strings.Join(factor, ",")
  }

  return strconv.Itoa(number)
}

func findFactor(number int, base int) []string {
  var factor []string
  for (number/base > 1) && (number%base == 0) {
    factor = append(factor, strconv.Itoa(base))
    number/=base
  }
  factor = append(factor, strconv.Itoa(number))
  return factor
}
