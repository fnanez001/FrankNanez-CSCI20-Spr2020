// Frank Nanez
// 4-15-20
// create a program that takes your total amount and the amount you wish to add in gratuity. Then presents the total plus the tip amount and the new total.

package main

import (
  "fmt"
  "math"
)

func tipA(totalA float64, grat float64) float64 {
  tip := (totalA*(grat*.01))
  return tip
}

func main() {
  var total float64
  var gratuity float64

  fmt.Println("This program will help you calculate the total with tip included based on your total and how much gratuity you would like to add. What what the total?")
  fmt.Scanln(&total)
  fmt.Println("Thank you and the gratuity amount you would like to add?")
  fmt.Scanln(&gratuity)

  tipAmount := tipA(total,gratuity)
  fmt.Println("So the total was,",total,"and the gratuity amount is,",gratuity,"%. The tip would be",(math.Round(tipAmount*100)/100),". Your grand total would be $",(math.Round((total+tipAmount)*100)/100),".")
}