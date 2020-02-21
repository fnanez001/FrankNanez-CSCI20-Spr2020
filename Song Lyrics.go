// Frank Nanez
// 2-20-20
// have song lyrics loop to how many times the user would like

package main

import "fmt"
//set variables
var (
  i = 0
  times = 0
)
// set variable for song lyric to loop
var songlyrics = "\"Hey, hey, good looking, whatcha got cooking?\""
// set reason for code, share lyrics and how many times they would like to have it repeated
func main() {
  fmt.Println("So this is your favorite part of your favorite song?")
  fmt.Println("\"Hey, hey, good looking, whatcha got cooking?\"")
  fmt.Println("How many times would you like to hear this?")
 // scan for number of times
  fmt.Scanln(&times)
 //set conditions for variable to loop = to the number of times requested
  for i = 0; i <= times; i++{
    fmt.Println(songlyrics)
  }
}