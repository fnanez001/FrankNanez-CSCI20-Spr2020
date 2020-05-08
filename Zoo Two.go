// Frank Nanez
// 4-30-20
// create a program to log zoo attendants and they work load. Also keeps track of animals 

package main

import "fmt"

type zooAttendant struct{
  Name string
  zAnimal1 string
  zAnimal2 string
  zAnimal3 string
  Zone string
}

type zZone struct{
  Name string
  Nofanimals int
  Cleaner string
}

type zAnimal struct{
  Name string
  Species string
  Age int
  Breakfast string
  Dinner string
}

func main() {
 var attend[3]zooAttendant
 attend[0] = zooAttendant{
   Name: "John",
   zAnimal1: "Izzy",// iguana
   zAnimal2: "Cam",//crocodile
   zAnimal3: "Tom",//turtle
   Zone: "reptiles",
 } 
 attend[1] = zooAttendant{
   Name: "Zoey",
   zAnimal1: "Pat",//Parrot
   zAnimal2: "Chica",//Crane
   zAnimal3: "Frankie",//finch
   Zone: "Birds",
 }
 attend[2] = zooAttendant{
   Name: "Mitch",
   zAnimal1: "Otto",// otter
   zAnimal2: "Edgar",// Elephant
   zAnimal3: "Cassie",//kangaroo
   Zone: "Mammals",
 }
 var zoneZ[3]zZone
 zoneZ[0] = zZone{
   Name: "Reptiles",
   Nofanimals: 3,
   Cleaner: "John",
 }
 zoneZ[1] = zZone{
   Name: "Birds",
   Nofanimals: 3,
   Cleaner: "Zoey",
 }
 zoneZ[2] = zZone{
   Name: "Mammals",
   Nofanimals: 3,
   Cleaner: "Mitch",
 }

 var zooA[9]zAnimal
 zooA[0] = zAnimal{
   Name: "Izzy",
   Species: "Reptile",
   Age: 4,
   Breakfast: "6am",
   Dinner: "4pm",
 }
 zooA[1] = zAnimal{
   Name: "Cam",
   Species: "reptile",
   Age: 12,
   Breakfast: "5am",
   Dinner: "5pm",
 }
 zooA[2] = zAnimal{
   Name: "Turtle",
   Species: "Reptile",
   Age: 20,
   Breakfast: "8am",
   Dinner: "6pm",
 }
 zooA[3] = zAnimal{
   Name: "Pat",
   Species: "Birds",
   Age: 5,
   Breakfast: "6am",
   Dinner: "6pm",
 }
 zooA[4] = zAnimal{
   Name: "Chica",
   Species: "Birds",
   Age: 2,
   Breakfast: "5:30am",
   Dinner: "5:30pm",
 }
 zooA[5] = zAnimal{
   Name: "Frankie",
   Species: "Birds",
   Age: 6,
   Breakfast: "6:30am",
   Dinner: "6:30pm",
 }
 zooA[6] = zAnimal{
   Name: "Otto",
   Species: "Mammals",
   Age: 4,
   Breakfast: "10am",
   Dinner: "6pm",
 }
 zooA[7] = zAnimal{
   Name: "Edgar",
   Species: "Mammals",
   Age: 15,
   Breakfast: "9am",
   Dinner: "5pm",
 }
 zooA[8] = zAnimal{
   Name: "Cassie",
   Species: "Mammals",
   Age: 7,
   Breakfast: "8am",
   Dinner: "4pm",
 }


 fmt.Println("Welcome to the Boston Zoo!")
 fmt.Println("We have 3 main zones for you to check out:")
 for i:=0; i<3; i++{
 fmt.Println(zoneZ[i].Name)
 }
 fmt.Println("Here are the animal names, zones, and feeding times for you to checkout out the exhibits:")
 for i:=0; i<9; i++{
   fmt.Println("Name:",zooA[i].Name,"Zone:",zooA[i].Species,"Feeding times:",zooA[i].Breakfast,"&",zooA[i].Dinner)
 }
 fmt.Println("For your convience the attendants for those areas are:")
 for i:=0; i<3; i++{
   fmt.Println(zoneZ[i].Name,"Attendant:",zoneZ[i].Cleaner)
 }
}