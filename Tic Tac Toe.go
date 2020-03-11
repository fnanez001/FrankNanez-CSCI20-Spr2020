// Frank Nanez
// 3-3-20
// play a tic tac toe game with user

//load package and import packages to program
package main

import "fmt"

func main() {
  //set array(s) for game
  var board [9]string
  board[0]= "o"
  board[1]= "o"
  board[2]= "o"
  board[3]= "o"
  board[4]= "o"
  board[5]= "o"
  board[6]= "o"
  board[7]= "o"
  board[8]= "o"
  
  //set var for user to use for input
  var userRune rune

   //Print for user to see board and explain game then ask for input
  fmt.Println( board[0],"|",board[1],"|",board[2],"\n",board[3],"|",board[4],"|",board[5],"\n",board[6],"|",board[7],"|",board[8])
  fmt.Println("Please select where you would like to enter an 'x' on the grid, starting in the top left acrosse then next line same thing and so on for the last line. i.e. top left is box 1 and top right is box 3. Then middle row left is 4.")

 // set loop for total turns of board
  for i := 0; i <= 9; i++ {
    //take input inside loop to repeat input
    fmt.Scanln(&userRune)
   // set inputs to board values
   if board[userRune-1] == "o" {
     board[userRune-1]=  "x"
     }else if board[userRune-1] != "o" {
       fmt.Println(" That has already been selected, pleae select another placement.")
    }
   fmt.Println( board[0],"|",board[1],"|",board[2],"\n",board[3],"|",board[4],"|",board[5],"\n",board[6],"|",board[7],"|",board[8])
  }
  // set condition to end game
  for j := 0; j <= 9; j++ {
    if board[j] != "o"{
    }
   
  }
  fmt.Println("Game is over!")
}