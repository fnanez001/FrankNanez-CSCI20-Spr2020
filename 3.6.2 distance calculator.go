// Frank Nanez
// 4-15-13
// create a program that takes coordinate and calculates the distance between them

package main

import (
  "fmt"
  "math"
)

func dist(x1 float64,x2 float64,y1 float64,y2 float64) (float64) {
  var distA = math.Pow((x2-x1),2)
  var distB = math.Pow((y2-y1),2)
  var distC = math.Sqrt(distA+distB)
 return distC
}

func main () {
  var xA,xB,yA,yB float64

  fmt.Println("This program will help you calculate your distance between two locations rounded to the closest whole mile. Please give the first coordinates as 'x,y'.")
  fmt.Scan(&xA,&yA)
  fmt.Println("Thank you and the second coordinates?'x,y'")
  fmt.Scan(&xB,&yB)

  distance := dist(xA,xB,yA,yB)

  fmt.Println("The total distance between those points is:",math.Round(distance),"miles")
}