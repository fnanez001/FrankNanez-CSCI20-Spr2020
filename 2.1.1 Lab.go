// Frank Nanez
// 2-18-20
// filling in code for comments provided.

package main

import (
    "fmt"
    "math/rand"
) //adding the ability to do random numbers

func main() {
    //create a variable for count
    var i = 0
    var max = 0
    var guess = 0
    //ask the user to enter a max range for the guessing game and store that value in variable max.
    fmt.Scanln(&max)

    //this next line creates a random number from 1 to that guess for the computer to know.  You can test this by printing out the variable computerGuess
    var computerGuess = rand.Intn(max)
    
    //ask the user to enter a guess for the computer number
    fmt.Scanln(&guess)

    //create a loop that compares the computerGuess to the userGuess while they are NOT equal go into the loop
        //increase the count by 1
        //tell the user that the guessed incorrect
        //ask the user to enter a new guess for the computer number
        for guess != computerGuess {
          i =i+1
          fmt.Println("incorrect")
          fmt.Println("Please enter a new guess")
          fmt.Scan(&guess)
        }
    
    //print out that the user got the answer correctly and how many guesses it took (the count)
    fmt.Println(i, "number of guesses to correct answer.")
}