// Programmer name: __________
// Date completed:  __________
// Description: ___________________________

package main

import (
    "fmt"
    "math/rand"
    "time"
) //adding the ability to do random numbers

func main() {
    //create two variables - one for the computer and one for the user
    var  user int

    rand.Seed(time.Now().UnixNano())
    //use a random integer value representing the computer's choice in a game of Rock, Scissors, Paper. 0=rock, 1=scissors, 2=paper
    a := (rand.Intn(3))

    //prompt the user for an integer value representing the player's choice
    fmt.Println("Please select Rock, Scissors, or Paper; 0=Rock, 1=Scissors, 2=Paper")
    fmt.Scanln(&user)
    fmt.Println("")

    //Print out the values using the words rock, scissors, paper.  ie. "Computer chose rock and player chose paper"
    if a == 0 && user == 0 {
     fmt.Println("Computer chose and you both chose rock") 
    } else if a == 0 && user == 1 {
      fmt.Println("Computer Chose rock and you chose scissors")
    } else if a == 0 && user == 2 {
      fmt.Println("Computer chose rock and you chose paper")
    }
    if a == 1 && user == 0 {
      fmt.Println("Computer chose scissors and you chose rock")
    } else if a == 1 && user == 1 {
      fmt.Println("You and the computer chose scissors")
    } else if a == 1 && user == 2 {
      fmt.Println("Computer choser scissors and you chose paper")
    }
    if a == 2 && user == 0 {
      fmt.Println("Computer chose paper and you chose rock")
    } else if a == 2 && user == 1 {
      fmt.Println("Computer chose paper and you chose scissors")
    } else if a == 2 && user == 2 {
      fmt.Println("You and the computer both chose paper")
    }
    //You will need to use decisions for this

}