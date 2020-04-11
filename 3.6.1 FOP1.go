// Frank Nanez
// 4-10-20
// create a function that converts seconds to hours,minutes, and seconds

package main

import "fmt"

func hoursMinutesseconds(seconds int) (int, int, int) {
  hours := (seconds/60/60)
  minutes := (hours%60)/60
  second := (minutes%60)
	return hours, minutes, second
}

func main() {
  var time int
  var hours, minutes, seconds = hoursMinutesseconds(hours,minutes,second)
//call your function
 fmt.Println("Please give the number of seconds you wish to convert ot hours, minutes, seconds")
 fmt.Scanln(&time)

 fmt.Println("Hours:",hours,"Minutes:",minutes,"Seconds:",seconds)

}