// Frank Nanez
// 3-10-20
// play NIM with computer and user

// load your package and import the proper ones needed for the program
package main

import (
  "fmt"
   "math/rand"
   "time"
)

func main() {

//set a random seed for random numbers
rand.Seed(time.Now().UnixNano())

//set variables for piles and print the values for NIM
pile1 := (rand.Intn(19)+1)
pile2 := (rand.Intn(19)+1)
pile3 := (rand.Intn(19)+1)
fmt.Println("Pile 1:",pile1,"Pile 2:",pile2,"Pile 3:",pile3)

//set variables for user and computer
 var playerT int
 var pPile int
 var winner string 
 // explain game to player
 fmt.Println("In this game you play against the computer, removing 1,2, or 3 stones from each pile. Winner removes the last stones.")

 // set loop for piles to be greater than 0 so game loops until all piles are 0
for pile1 > 0 || pile2 > 0 || pile3 > 0 {
   //ask player for  amount and user input
   fmt.Println("Please enter your amount you wish to remove from the pile.")
   fmt.Scanln(&playerT)
   //ask user for which pile and input
   fmt.Println("Please select the pile you wish to remove the stones from. Select pile 1,2, or 3.")
   fmt.Scanln(&pPile)
   //set computer amount and pile
   cPile := (rand.Intn(2)+1)
   computerT := (rand.Intn(2)+1)
   // check piles to see if game is over after user turn and all piles are 0, if 0 after player turn then player wins
   if pPile == 1 {
     pile1 = (pile1-playerT)
   }else if pPile == 2 {
     pile2 = (pile2-playerT)
   }else if pPile == 3 {
     pile3 = (pile3-playerT)
   }
   //print player selection to help keep track of game
  fmt.Println("User picked: ", playerT,"\n", "Pile 1:", pile1, "Pile 2:", pile2, "Pile 3:", pile3)
   //check piles to see if game is over after user turn and all piles are 0, if 0 after player turn then computer wins
   if pile1 == 0 || pile2 == 0 || pile3 == 0 {
     winner = "Player"
   }else if cPile == 1 {
     pile1 = (pile1-computerT)
   }else if cPile == 2 {
     pile2 = (pile2-computerT)
   }else if cPile == 3 {
     pile3 = (pile3-computerT)
   }else {
     winner = "Computer"
   }
   //print computer selection to help keep track of game 
     fmt.Println("Computer picked: ", computerT, "Pile :",cPile)
   fmt.Println("Pile 1:",pile1,"Pile 2:",pile2,"Pile 3:",pile3)
   
  }
  //once loop has ended then game is over, anounce the winner
  fmt.Println("Winner is",winner,"!")
}