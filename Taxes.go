// Frank Nanez
// 2-25-20
// have user inter their status and income to give them your their tax bracket and owed amount

package main

import "fmt" 

func main() {
  // define the variables for the income and filing status
  var income float32
  var status string

 // Prompt user variables needed
  fmt.Println("This program will determine your taxes owed and the tax bracket based on your filing status and income")
  fmt.Println("Please enter the corresponding letter to your filing status; A-Single, B-Marries filing joingtly, C-Head of household, D-Married filing separately")

  // have user enter their status letter and income
  fmt.Scanln(&status)

  fmt.Println("Thank you, please enter your gross income next.")

  fmt.Scanln(&income)

 // create if statements for status and else if for income brackets and amounts owed.
  if status == "A" {
     fmt.Println("You selected the filing status Single.")
     }else if income >= 0 && income <= 9700 && status != "B" && status != "C" && status != "D"{
     fmt.Println("$",income*.1,"Is the amount owed and you would be in the 10% tax bracket.")
     }else if income >= 9701 && income <= 39475 && status != "B" && status != "C" && status != "D"{
     fmt.Println("$",income*.12,"Is the amount owed and you would be in the 12% bracket")
     }else if income >= 39476 && income <= 84200 && status != "B" && status != "C" && status != "D"{
     fmt.Println("$",income*.22,"Is the amount owed and you would be in the 22% tax bracket.")
     }else if income >= 84201 && income <= 160725 && status != "B" && status != "C" && status != "D"{
     fmt.Println("$",income*.24,"Is the amount owed and you would be in the 24% tax bracket.")
     }else if income >= 160726 && income <= 204100 && status != "B" && status != "C" && status != "D"{
     fmt.Println("$",income*.32,"is the amount owed and you would be in the 32% tax bracket.")
     }else if income >= 204101 && income <= 510300 && status != "B" && status != "C" && status != "D"{
     fmt.Println("$",income*35,"Is the amount owed and you would be in the 35% tax bracket.")
     }else if income >= 510301 && status != "B" && status != "C" && status != "D"{
     fmt.Println("$",income*.37,"Is the amount owed and you would be in the 37% tax bracket.")
     }
  
  if status == "B" {
     fmt.Println("You selected the filing status Married and filing jointly.")
     }else if income >= 0 && income <= 19400 && status != "A" && status != "C" && status != "D"{
     fmt.Println("$",income*.1,"Is the amount owed and you would be in the 10% tax bracket.")
     }else if income >= 19401 && income <= 78950 && status != "A" && status != "C" && status != "D"{
     fmt.Println("$",income*.12,"Is the amount owed and you would be in the 12% bracket")
     }else if income >= 78951 && income <= 168400 && status != "A" && status != "C" && status != "D"{
     fmt.Println("$",income*.22,"Is the amount owed and you would be in the 22% tax bracket.")
     }else if income >= 168401 && income <= 321450 && status != "A" && status != "C" && status != "D"{
     fmt.Println("$",income*.24,"Is the amount owed and you would be in the 24% tax bracket.")
     }else if income >= 321451 && income <= 408200 && status != "A" && status != "C" && status != "D"{
     fmt.Println("$",income*.32,"is the amount owed and you would be in the 32% tax bracket.")
     }else if income >= 408201 && income <= 612350 && status != "A" && status != "C" && status != "D"{
     fmt.Println("$",income*35,"Is the amount owed and you would be in the 35% tax bracket.")
     }else if income >= 612351 && status != "A" && status != "C" && status != "D"{
     fmt.Println("$",income*.37,"Is the amount owed and you would be in the 37% tax bracket.")
     }
   
  if status == "C" {
     fmt.Println("You selected the filing status Head of household.")
     }else if income >= 0 && income <= 13850 && status != "A" && status != "B" && status != "D"{
     fmt.Println("$",income*.1,"Is the amount owed and you would be in the 10% tax bracket.")
     }else if income >= 13851 && income <= 52850 && status != "A" && status != "B" && status != "D"{
     fmt.Println("$",income*.12,"Is the amount owed and you would be in the 12% bracket")
     }else if income >= 52851 && income <= 84200 && status != "A" && status != "B" && status != "D"{
     fmt.Println("$",income*.22,"Is the amount owed and you would be in the 22% tax bracket.")
     }else if income >= 84201 && income <= 160700 && status != "A" && status != "B" && status != "D"{
     fmt.Println("$",income*.24,"Is the amount owed and you would be in the 24% tax bracket.")
     }else if income >= 160701 && income <= 204100 && status != "A" && status != "B" && status != "D"{
     fmt.Println("$",income*.32,"is the amount owed and you would be in the 32% tax bracket.")
     }else if income >= 204101 && income <= 510300 && status != "A" && status != "B" && status != "D"{
     fmt.Println("$",income*35,"Is the amount owed and you would be in the 35% tax bracket.")
     }else if income >= 510301 && status != "A" && status != "B" && status != "D"{
     fmt.Println("$",income*.37,"Is the amount owed and you would be in the 37% tax bracket.")
     }
  
  if status == "D" {
     fmt.Println("You selected the filing status Married filing separately.")
     }else if income >= 0 && income <= 9700 && status != "A" && status != "B" && status != "C"{
     fmt.Println("$",income*.1,"Is the amount owed and you would be in the 10% tax bracket.")
     }else if income >= 9701 && income <= 39475 && status != "A" && status != "B" && status != "C"{
     fmt.Println("$",income*.12,"Is the amount owed and you would be in the 12% bracket")
     }else if income >= 39476 && income <= 84200 && status != "A" && status != "B" && status != "C"{
     fmt.Println("$",income*.22,"Is the amount owed and you would be in the 22% tax bracket.")
     }else if income >= 84201 && income <= 160725 && status != "A" && status != "B" && status != "C"{
     fmt.Println("$",income*.24,"Is the amount owed and you would be in the 24% tax bracket.")
     }else if income >= 160726 && income <= 204100 && status != "A" && status != "B" && status != "C"{
     fmt.Println("$",income*.32,"is the amount owed and you would be in the 32% tax bracket.")
     }else if income >= 204101 && income <= 306175 && status != "A" && status != "B" && status != "C"{
     fmt.Println("$",income*35,"Is the amount owed and you would be in the 35% tax bracket.")
     }else if income >= 306176 && status != "A" && status != "B" && status != "C"{
     fmt.Println("$",income*.37,"Is the amount owed and you would be in the 37% tax bracket.")
     }
  
}