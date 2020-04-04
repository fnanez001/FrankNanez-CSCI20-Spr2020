


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

  cAnswer := (rand.Intn(1)+1)
  i := cAnswer
  fmt.Println(correctAns[i])
}

func incorrect () {
  rand.Seed(time.Now().UnixNano())

  var incorrectAns [3]string
  incorrectAns[0] = "Close!"
  incorrectAns[1] = "Sorry that is incorrect, please try again."
  incorrectAns[2] = "Unfortunately that is not the answer."

  inAnswer := (rand.Intn(1)+1)
  i := inAnswer
  fmt.Println(incorrectAns[i])
}

func main () {
  
}