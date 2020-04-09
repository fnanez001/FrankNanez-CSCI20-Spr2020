// Frank Nanez
// 4-8-20
// take user score and prints out grade based on score

package main

import "fmt"

func grade (grade int){
  
  if grade >= 90{
    fmt.Println("Your grade is an 'A'.")
  } else if grade >= 80 && grade < 90{
    fmt.Println("Your grade is a 'B'.")
  } else if grade >= 70 && grade < 80{
    fmt.Println("Your grade is a 'C'.")
  } else if grade < 70{
    fmt.Println("Sorry you failed. Please try again.")
  }
}

func main () {
  var score int

  fmt.Println("Please enter your score and you grade will be presented.")
  fmt.Scanln(&score)

  grade(score)
}