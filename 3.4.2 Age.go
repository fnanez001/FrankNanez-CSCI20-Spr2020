// Frank Nanez
// 4-8-20
// program  to take user birth year and name and prints out users age using their name

package main

import (
  "fmt"
  "time"
)

var currentTime = time.Now()

func age (person string, age int) {
  
  fmt.Println("Oh cool",person,"so you're",currentTime.Year()-age,"years old.")
}

func main () {
  var name string
  var year int

  fmt.Println("Hello. What is your name?")
  fmt.Scanln(&name)
  fmt.Println("What year were you born?")
  fmt.Scanln(&year)
  age(name,year)
}
