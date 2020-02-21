// Frank Nanez
// 2-20-20
// have user enter score until the enter -1 and the n the average of all scores will be entered.

package main

import "fmt" 
//set variables needed for # of times, the score to be entered, and total of score
var (
  i = 0
  score = 0
  total = 0
)
//ask for scores and explain the 
func main(){
  fmt.Println("Please enter your test scores. After entering your final test score then you will be given the average of your scores.")
  fmt.Scanln(&score)
 // create loop; set variable, condition, and change condition for above set up
  for i = 0; score >= 0; i++ {
   total = total+score
   fmt.Println("Please enter your score")
   fmt.Scanln(&score)
  }
  // print average after loop is excuted
 fmt.Println(total/i)
}