// Frank Nanez
// 3-3-20
// play a tic tac toe game with user

package main

import "fmt"

func main() {
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
  
  var userRune rune

   
  fmt.Println( board[0],"|",board[1],"|",board[2],"\n",board[3],"|",board[4],"|",board[5],"\n",board[6],"|",board[7],"|",board[8])
  fmt.Println("Please select where you would like to enter an 'x' on the grid, starting in the top left acrosse then next line same thing and so on for the last line. i.e. top left is box 1 and top right is box 3. Then middle row left is 4.")

  for i := 0; i <= 9; i++ {
    fmt.Scanln(&userRune)

   if board[userRune-1] == "o" {
     board[userRune-1]=  "x"
     }else if board[userRune-1] != "o" {
       fmt.Println(" That has already been selected, pleae select another placement.")
    }
   fmt.Println( board[0],"|",board[1],"|",board[2],"\n",board[3],"|",board[4],"|",board[5],"\n",board[6],"|",board[7],"|",board[8])
  }
  for j := 0; j <= 9; j++ {
    if board[j] != "o"{
    }
   
  }
  fmt.Println("Game is over!")
}