package raindrops

import "strconv"

func Convert(r int) string {
  if r%105 == 0 {
    return "PlingPlangPlong"
  } else if r%35 == 0 {
    return "PlangPlong"
  } else if r%21 == 0 {
    return "PlingPlong"
  } else if r%15 == 0 {
    return "PlingPlang"
  } else if r%7 == 0 {
    return "Plong"
  } else if r%5 == 0 {
    return "Plang"
  } else if r%3==0 {
    return "Pling"
  }
  return strconv.Itoa(r)
  }
