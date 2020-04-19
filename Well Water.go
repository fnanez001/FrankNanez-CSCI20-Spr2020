// Frank Nanez
// 4-16-20
// create a program that takes radius of a well and total depth in feet and calculate the total gallons of the well.

// set up your packages and import necessary packages
package main

import (
  "fmt"
  "math"
)

// create a constant to use for Pi in your area of a circle
const Pi  = 3.14159265358979323846264338327950288419716939937510582097494459

// set up a function for finding a the cubic ft of the well (minus) the water depth if possible to get an accurate amount of gallons in the end.
func cubicF( rad float64, dep float64, wDep float64) float64 {
  radInft := rad/12
  area := (Pi*(radInft*radInft))
  wDepth := dep-wDep
  cubic := area*wDepth
  return cubic
}

// create a function to convert the cubic feet to gallons
func convertGal(cub float64) float64 {
  gallons := cub*7.481
  return gallons
}

func enoughW (gal float64) float64 {

  if gal >= 250{
    gal = 1
  }else if gal < 250{
    gal = 2
  }
  
  return gal
}

// use main func to prompt user for radius of well, depth, and the depth down to the water table. then call functions you created above to get the accurate amount of gallons for users well
func main() {
  var radius float64
  var depth float64
  var waterD float64

  fmt.Println("This program will help you calculate the gallons of your well, or the gallons of a well you plan to install, based on the radius[in inches] of your well and the depth[in feet] to see if it will be sufficient for a mily home. We will also need the depth down to the water table to help calculate accurately. First, please enter your radius of the well.")
  fmt.Scanln(&radius)
  fmt.Println("Thank you. Second, we will need the depth in feet.")
  fmt.Scanln(&depth)
  fmt.Println("Great! Lastly we need the depth down to the water table in feet.")
  fmt.Scanln(&waterD)

  wCubft := cubicF(radius,depth,waterD)
  totalGal := convertGal(wCubft)
  enough := enoughW(totalGal)


  fmt.Println("So the total gallons of the spec'd well would be",math.Round(totalGal*100)/100)
  
  if enough == 1{
    fmt.Println("This would be enough water for the needs of a family of 4.")
  }else {
    fmt.Println("This would not cover the needs of a family of 4.")
  }
 

}