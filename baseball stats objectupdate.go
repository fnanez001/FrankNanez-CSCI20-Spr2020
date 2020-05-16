// Frank Nanez
// 3-3-20
// giving player name, batting averages, and slugging averages and the highest and lowest scores 

// import packages needed
package main

import "fmt"

// create struct for players. Includes name, batting average, slugging average
type player struct {
  Name string
  Batting float64
  Slugging float64
}

func main() {

  // create array for player struct to include 4 players and their stats.
  var bPlayer[4]player
  bPlayer[0] = player{
    Name: "Joe",
    Batting: 1.2,
    Slugging: 6.1,
  }
  bPlayer[1] = player{
    Name: "Bill",
    Batting: 2.1,
    Slugging: .99,
  } 
  bPlayer[2] = player{
     Name: "Frank",
     Batting: 4.5,
     Slugging: 5.1,
  }
  bPlayer[3] = player{
     Name: "JaShen",
     Batting: 1.1,
     Slugging: 1.0,
  }
  
   // prints statement with players stats 
    fmt.Println("Here are the player stats: Name, Batting Average, Slugging Average")
  for i := 0; i < 4; i++ {
  fmt.Println(bPlayer[i].Name,",",bPlayer[i].Batting,",", bPlayer[i].Slugging)
  }
 
 // create variables to use to get highest batting average
  highestB := bPlayer[0].Batting
  highestI := 0
 // loop to search for highest batting average 
  for i := 0; i <= 3; i++ {
    if bPlayer[i].Batting > highestB {
    highestB = bPlayer[i].Batting 
    highestI = i
    }
    
  }
  // print results for highest batting average
  fmt.Println("Highest batting average: ",bPlayer[highestI])
  
  // create new variable to use to get highest slugging average
  highestS := bPlayer[0].Slugging
  highestI = 0

  // loop to search for highest slugging average
  for i := 0; i <= 3; i++ {
    if bPlayer[i].Slugging > highestS {
    highestS = bPlayer[i].Slugging
    highestI = i
    }

  }
  // print results for highest slugging average 
  fmt.Println("Highest slugging average: ",bPlayer[highestI])
  
  //create new variable to use to get lowest batting average
  lowestB := bPlayer[0].Batting
  lowestI := 0

  // loop to search for lowest batting average
  for i := 0; i <= 3; i++ {
    if bPlayer[i].Batting < lowestB {
      lowestB = bPlayer[i].Batting
      lowestI = i
    }

  }
  // print results for lowest batting average
  fmt.Println("Lowest batting average:",bPlayer[lowestI])

  // create a new varaible to use to get lowest slugging average
  lowestS := bPlayer[0].Slugging
  lowestI = 0

  // loop to search for lowest slugging average
  for i := 0; i <= 3; i++ {
   if bPlayer[i].Slugging < lowestS {
     lowestS = bPlayer[i].Slugging
     lowestI = i
   }

  }
  // print results for lowest slugging average
  fmt.Println("Lowest slugging average:",bPlayer[lowestI])

}