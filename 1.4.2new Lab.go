// Frank Nanez
// 2-6-20
// Canoe project to convert rods into meters, feet, and miles for canoer.

package main

import "fmt" 

//set variable for starting point
func main() {
  var distanceInrods = 0.0

  //Set up welcome and question for user to select.
  fmt.Println("Welcome canoefarer! Please select a path you wish to know about.The information display for each trail will be the length in rod(s) given,the length in feet, the length in meter(s), the length in mile(s), and of course the average time to travel the distance on foot with your canoe. Measurements are shown in rods, please give the respective rods measurement you wish to learn about")

  fmt.Println("NOTE: Trail C is the same distance to the starting point at the parking lot to each of the lakes.")

  // give visual so the user knows how to respond.
  fmt.Println("Trail A is 37.5 rods from Lake A to lake B")
  fmt.Println("Trail B is 55.2 rods from Lake B to lake C")
  fmt.Println("Trail C is 30 rods back to the parking lot from any of the lakes")

  //ask for user input on distance in rods to get selected trail.
  fmt.Println("Please input your desired trail but inputing the amount of rods the trail is.")
  fmt.Scanln(&distanceInrods)

  fmt.Println(distanceInrods,distanceInrods*16.5,"feet", (distanceInrods*16.5)/3.0,"meters", (distanceInrods*16.5)/5280.0,"miles", "average travel time would be", (((distanceInrods*16.5)/5280.0)*3.1),"hours")

  

}
