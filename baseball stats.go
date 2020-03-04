// Frank Nanez
// 3-3-20
// giving player name, batting averages, and slugging averages and the highest and lowest scores 

package main

import "fmt"

func main() {
  var player = [4] string {
    "Joe",
    "Bill",
    "Frank",
    "JaShen",
    }

  var battingAvg = [4] float64 {
    1.2,
    2.1,
    4.5,
    1.1,
    }
  
  var sluggingAvg = [4] float64 {
    6.1,
    .99,
    5.1,
    1.0,
    }

    fmt.Println("Here are the player stats: Name, Batting Average, Slugging Average")
  for i := 0; i < 4; i++ {
  fmt.Println(player[i],",", battingAvg[i],",", sluggingAvg[i])
  }

  highestB := battingAvg[0]
  highestI := 0

  for i := 0; i <= 3; i++ {
    if battingAvg[i] > highestB {
    highestB = battingAvg[i] 
    highestI = i
    }
    
  }
  fmt.Println("Highest batting average: ",player[highestI],",", battingAvg[highestI],",", sluggingAvg[highestI])

  highestS := sluggingAvg[0]
  highestI = 0

  for i := 0; i <= 3; i++ {
    if sluggingAvg[i] > highestS {
    highestS = sluggingAvg[i]
    highestI = i
    }

  }
  fmt.Println("Highest slugging average: ",player[highestI],",", battingAvg[highestI],",", sluggingAvg[highestI])

  lowestB := battingAvg[0]
  lowestI := 0

  for i := 0; i <= 3; i++ {
    if battingAvg[i] < lowestB {
      lowestB = battingAvg[i]
      lowestI = i
    }

  }
  fmt.Println("Lowest batting average:",player[lowestI],",",battingAvg[lowestI],",", sluggingAvg[lowestI])

  lowestS := sluggingAvg[0]
  lowestI = 0

  for i := 0; i <= 3; i++ {
   if sluggingAvg[i] < lowestS {
     lowestS = sluggingAvg[i]
     lowestI = i
   }

  }
  fmt.Println("Lowest slugging average:",player[lowestI],",",battingAvg[lowestI],",", sluggingAvg[lowestI])

}