// Frank Nanez
// 2-20-20
// have song lyrics loop to how many times the user would like

package main

import "fmt"

var (
  i = 0
  times = 0
)
var songlyrics = "\"Hey, hey, good looking, whatcha got cooking?\""

func main() {
  fmt.Println("So this is your favorite part of your favorite song?")
  fmt.Println("\"Hey, hey, good looking, whatcha got cooking?\"")
  fmt.Println("How many times would you like to hear this?")

  fmt.Scanln(&times)

  for i = 0; i <= times; i++{
    fmt.Println(songlyrics)
  }
}