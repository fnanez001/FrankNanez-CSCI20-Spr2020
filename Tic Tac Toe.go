// Frank Nanez
// 3-3-20
// play a tic tac toe game with user

package main

import (
   "fmt"
   "math/rand"
   "time"
)

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
  rand.Seed(time.Now().UnixNano())
  
   
  fmt.Println( board[0],"|",board[1],"|",board[2],"\n",board[3],"|",board[4],"|",board[5],"\n",board[6],"|",board[7],"|",board[8])
  fmt.Println("Please select where you would like to enter an 'x' on the grid, starting in the top left acrosse then next line same thing and so on for the last line. i.e. top left is box 1 and top right is box 3. Then middle row left is 4.")
  fmt.Scanln(&userRune)
  computer := (rand.Intn(9))
  
  for i := 0; i < 10; i++ {
   if board[userRune-1] == "o" && board[computer] == "o" {
     board[userRune-1]=  "x"
     board[computer] = "0"
     fmt.Println( board[0],"|",board[1],"|",board[2],"\n",board[3],"|",board[4],"|",board[5],"\n",board[6],"|",board[7],"|",board[8])
    }else if board[computer] == board[userRune-1] {
     computer = (rand.Intn(9))
     board[computer] = "0"
    }else if board[userRune-1] != "o" && board[computer] != "o"{
     fmt.Println(" That has already been selected, pleae select another placement.")
    }
    fmt.Println("Where would you like to place an 'x' next?")
    fmt.Scanln(&userRune)
  }
  
  for j := 0; j < 10; j++ {
    if board[j] != "o"{
     fmt.Println("Game is over!")
     }
   }

}