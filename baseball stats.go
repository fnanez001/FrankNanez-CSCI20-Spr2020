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

  var battingAvg = [4] float32 {
    1.2,
    2.1,
    4.5,
    1.1,
    }
  
  var sluggingAvg = [4] float32 {
    1.1,
    4.3,
    5.1,
    1.0,
    }

    fmt.Println("Here are the player stats: Name, Batting Average, Slugging Average")
  for i := 0; i < 4; i++ {
  fmt.Println(player[i],",", battingAvg[i],",", sluggingAvg[i])
  }

  highestB := battingAvg[0]
  
  for i := 0; i <+ 3; i++ {
    battingAvg[i] > highestB
    highestB = battingAvg[i]
    fmt.Println(highestB[i])
  }
 

}