// Frank Nanez
// 4-16-20
// create a program that takes radius of a well and total depth in feet and calculate the total gallons of the well.

package main

import (
  "fmt"
  "math"
)

const Pi  = 3.14159265358979323846264338327950288419716939937510582097494459


func cubicF( rad float64, dep float64, wDep float64) float64 {
  radInft := rad/12
  area := (Pi*(radInft*radInft))
  wDepth := dep-wDep
  cubic := area*wDepth
  return cubic
}

func convertGal(cub float64) float64 {
  gallons := cub*7.481
  return gallons
}

func main() {
  var radius float64
  var depth float64
  var waterD float64

  fmt.Println("This program will help you calculate the gallons of your well, or the gallons of a well you plan to install, based on the radius[in inches] of your well and the depth[in feet]. We will also need the depth down to the water table to help calculate accurately. First, please enter your radius of the well.")
  fmt.Scanln(&radius)
  fmt.Println("Thank you. Second, we will need the depth in feet.")
  fmt.Scanln(&depth)
  fmt.Println("Great! Lastly we need the depth down to the water table in feet.")
  fmt.Scanln(&waterD)

  wCubft := cubicF(radius,depth,waterD)
  totalGal := convertGal(wCubft)

  fmt.Println("So the total gallons of the spec'd well would be",math.Round(totalGal*100)/100)
}