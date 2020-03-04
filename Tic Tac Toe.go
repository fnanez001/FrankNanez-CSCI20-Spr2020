// Frank Nanez
// 3-3-20
// play a tic tac toe game with user

package main

import "fmt"

func main() {
  var board [10]string
  board[0]= "o"
  board[1]= "o"
  board[2]= "o"
  board[3]= "o"
  board[4]= "o"
  board[5]= "o"
  board[6]= "o"
  board[7]= "o"
  board[8]= "o"
  board[9]= "o"
  
  var userRune rune
   
  fmt.Println( board[0],"|",board[1],"|",board[2],"\n",board[3],"|",board[4],"|",board[6],"\n",board[7],"|",board[8],"|",board[9])
  fmt.Println("Please select where you would like to enter an 'x' on the grid, starting in the top left acrosse then next line same thing and so on for the last line. i.e. top left is box 1 and top right is box 3. Then middle row left is 4.")
  fmt.Scanln(&userRune)
  
}