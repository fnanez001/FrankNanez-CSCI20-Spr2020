// Frank Nanez
// 2-13-20
// echange rate and expected change rate in 1 yr

package main

import "fmt"

func main() {
  var dollars = 0.0
 //set up the situation and explain the conversion rate
  fmt.Println("Current exchange rate is .8:1 for the US dollar to the Euro. In 1 year it is expected that the value of the euro will rise 8% and the dollar will fall 10% of their current values. Please enter your current amount you wish to see what the conversion to euros would be.")

  //ask for dollar amount input
  fmt.Scanln(&dollars,)
  fmt.Println(dollars,"dollars")

 // convert dollar amount to euros for current conversion rate
  fmt.Println(dollars*.8,"euros")

 //Explain and print what that conversion would be in one year
  fmt.Println("In one year that same amount would roughly estimate to in euros based on the expected conversion rates changes.")

  //use the base euro conversion then substract the difference in the new rate for the new total.
  fmt.Println(dollars*.8-((dollars*.08)*.72),"euros")
}