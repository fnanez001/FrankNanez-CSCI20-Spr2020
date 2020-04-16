// Frank Nanez 
// 4-15-20
// create a program that takes a total amount in sale and determines by bracket what the commission amount is

// load package and import necessary items needed for program
package main

import "fmt"

// create a function that takes sales and calculates commission based on brackets for 4%, 6%, and 6.25% brackets given and return commission
func commissions(amount float64) float64 {
  var com float64

  if amount < 50000{
    com = (amount*.04)
  }else if amount >= 50000 && amount <= 250000 {
    com = (amount*.06)
  }else if amount > 250000 {
    com = (amount*.0625) 
  } 
 fmt.Println(com)
 return com
}

//create main function and ask for sales to use for commission, then call commission function to calculate and prints necessary staetments.
func main () {
  var sales float64

  fmt.Println("This will help you calculate your commisions based on your sales numbers. Please enter your total sales in whole dollars.")
  fmt.Scanln(&sales)

  commission := commissions(sales)
  fmt.Println("Great, you sold: $",sales,"and your commission would be: $",commission,"!")

}