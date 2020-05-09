// Frank Nanez
// 4-30-20
// create a struct for the some planets and calculate the weight based on the gravity of the planet.

package main

import (
  "fmt"
  "strings"
)

type Planet struct{
  Name string
  Gravity float64
  Moons int
}

func weightOnplanet(gravity float64, weight float64) (Weight float64){
 Weight = gravity*weight 
 return Weight
}

func main(){
  var planet[8]Planet
  planet[0] = Planet{
    Name: "Mecury",
    Gravity: 0.378,
    Moons: 0,
  }
  planet[1] = Planet{
    Name: "Venus",
    Gravity: 0.907,
    Moons: 0,
  }
  planet[2] = Planet{
    Name: "Earth",
    Gravity: 1.0,
    Moons: 1,
  }
  planet[3] = Planet{
    Name: "Mars",
    Gravity: 0.377,
    Moons: 2,
  }
  planet[4] = Planet{
    Name: "Jupiter",
    Gravity: 2.36,
    Moons: 79,
  }
  planet[5] = Planet{
    Name: "Saturn",
    Gravity: 0.916,
    Moons: 82,
  }
  planet[6] = Planet{
    Name: "Uranus",
    Gravity: 0.889,
    Moons: 27,
  }
  planet[7] = Planet{
    Name: "Neptune",
    Gravity: 1.12,
    Moons: 14,
  }

  var weight float64
  for i:=0; i<8; i++{
    fmt.Println(planet[i].Name,"gravity:",planet[i].Gravity,"number of moons:",planet[i].Moons)
  }

  var pName string
  fmt.Println("Above are the planets, the gravity in comparison to Earth's gravity, and the number of moons for that planet. If you select a planet and enter your weight we can tell you what your wieght would be on those planets. Let's start with the name of the planet you would like to select?")
  fmt.Scanln(&pName)
  fmt.Println("Great! Next is your weight, in pounds, here on Earth?")
  fmt.Scanln(&weight)
  pName = strings.Title(pName)

  x := 0
  for i:=0; i<8; i++{
    if pName == planet[i].Name{
     fmt.Println("the gravity on that planet is",planet[x].Gravity,"Compared to Earth's.")
     newWeight := weightOnplanet(planet[x].Gravity,weight)
     fmt.Println("Your weight would be:",newWeight,"lbs.")
      i = 9
    }
    x=x+1
  }
  


}