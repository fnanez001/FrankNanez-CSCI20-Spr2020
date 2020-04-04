


package main

import (
   "fmt"
   "math/rand"
   "time"
)

func correct () {
 rand.Seed(time.Now().UnixNano())

  var correctAns [3]string
  correctAns[0] = "Great job!"
  correctAns[1] = "Nice!"
  correctAns[2] = "Nailed it!"

  cAnswer := (rand.Intn(3))
  i := cAnswer
  fmt.Println(correctAns[i])
}

func incorrect () {
  rand.Seed(time.Now().UnixNano())

  var incorrectAns [3]string
  incorrectAns[0] = "Close!"
  incorrectAns[1] = "Sorry that is incorrect, please try again."
  incorrectAns[2] = "Unfortunately that is not the answer."

  inAnswer := (rand.Intn(3))
  i := inAnswer
  fmt.Println(incorrectAns[i])
}

func main () {

  var userAns1 string

  fmt.Println("Today we will be quizzing on you something from your CSCI 20")
  fmt.Println("Question 1: In programming what is a variable?", "\n", "A: a location in memeory where data is stored", "\n", "B: a name of something you want to call", "\n", "C: something that can change at anytime")
  fmt.Scanln(&userAns1)

  if userAns1 == "A" || userAns1 == "a" {
    correct()
  }else if userAns1 !="A" || userAns1 != "a" {
    incorrect()
  }
}