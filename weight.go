package main

import (
  "fmt"
)

func main() {
  var weight int
  var weightone, weighttwo, weightthree, weightfour, weightfive int

  weightone = 175
  weighttwo = 160
  weightthree = 60
  weightfour = 200
  weightfive = 90

  weight = (weightone + weighttwo + weightthree + weightfour + weightfive)/5
  fmt.Print("average weight:", weight)

}
