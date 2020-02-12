// Frank Nanez
// 2-6-20
// Canoe project to convert rods into meters, feet, and miles for canoer.

package main

import "fmt" 

func main() {
  //set variable for starting point
  var distanceInrods = float32

  //Set up welcome and question for user to select.
  fmt.Println("Welcome canoefarer! Please select a path you wish to know about.The information display for each trail will be the length in rod(s) given, the length in meter(s), the length in mile(s), and of course the average time to travel the distance on foot with your canoe. Measurements are shown in rods, please give the respective rods measurement you wish to learn about")
  fmt.Println("NOTE: Trail C is the same distance to the starting point at the parking lot to each of the lakes.")

  // give visual so the user knows how to respond.
  fmt.Println("             ^         ")
  fmt.Println(" ^          / \        ^ ")
  fmt.Println("/ \  trl A  | | trl B / \")
  fmt.Println("| |------>  \ /-----> | |")
  fmt.Println("\ /  37.5    v  55.2  \ /")
  fmt.Println("   \         |         / ")
  fmt.Println("    \        |        /  ")
  fmt.Println("      >      v       <   ")
  fmt.Println("            trl C        ")

  //ask for user input on distance in rods to get selected trail.
  fmt.Println("Please input your desired trail but inputing the amount of rods the trail is.")
  fmt.Scanln(&distanceInrods)

}
