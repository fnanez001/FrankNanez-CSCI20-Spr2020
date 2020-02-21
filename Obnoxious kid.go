// Frank Nanez
// 2-20-20
// have kid ask are we there yet the number of times equal to their age

package main

import "fmt" 
// set variables
var (
  i = 0
  age = 0
)
// set scenario and ask how old the child is to have them loop are we there yet that many times.
func main() {
  fmt.Println("Child: How long is this trip going to take")
  fmt.Println("Driver: We will get there when we get there.")
  fmt.Println("Child: Are we there yet?")
  fmt.Println("Driver: Please don't ask again")
  fmt.Println("Driver: I got it, how old are you?")
 // scan for age and repeat back to child
  fmt.Scanln(&age)
  fmt.Println("Driver: Nice! So you are",age,"!")
 // set scenario for loop to equal that of age 
  for i = 0; i <= age; i++{
    fmt.Println("Kid: Are we there yes?!?!")
}

}