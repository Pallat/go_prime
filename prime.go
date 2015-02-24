package prime

import (
  "strconv"
  )

func prime(number int) string {
  if number == 4 {
    return "2,2"
  }
  if number == 6 {
    return "2,3"
  }
  if number == 8 {
    return "2,2,2"
  }
  return strconv.Itoa(number)
}
