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
  fmt.Println("Please enter the corresponding letter to your filing status; A-Single, B-Married filing joingtly, C-Head of household, D-Married filing separately")

  // have user enter their status letter and income
  fmt.Scanln(&status)

  fmt.Println("Thank you, please enter your gross income next.")

  fmt.Scanln(&income)

 // create if statements for status and else if for income brackets and amounts owed.
  if status == "A" && status != "B" && status != "C" && status != "D" {
     fmt.Println("You selected the filing status Single.")
     if income >= 0 && income <= 9700 {
     fmt.Println("$",income*.1,"Is the amount owed and you would be in the 10% tax bracket.")
     }else if income >= 9701 && income <= 39475 {
     fmt.Println("$",((income-9700)*.12+income*.1),"Is the amount owed and you would be in the 12% bracket")
     }else if income >= 39476 && income <= 84200 {
     fmt.Println("$",((income-39475)*.22)+(39475-9700)*.12+(9700*.1),"Is the amount owed and you would be in the 22% tax bracket.")
     }else if income >= 84201 && income <= 160725 {
     fmt.Println("$",(((income-84200)*.24)+(84200-39475)*.22+((39475-9700)*.12)+((9700)*.1)),"Is the amount owed and you would be in the 24% tax bracket.")
     }else if income >= 160726 && income <= 204100 {
     fmt.Println("$",(income-160725)*.32+(160725-84200)*.24+(84200-39475)*.22+(39475-9700)*.12+(9700*.1),"Is the amount owed and you would be in the 32% tax bracket.")
     }else if income >= 204101 && income <= 510300 {
     fmt.Println("$",(income-204100)*35+(204100-160725)*.32+(160725-84200)*.24+(84200-39475)*.22+(39475-9700)*.12+(9700*.1),"Is the amount owed and you would be in the 35% tax bracket.")
     }else if income >= 510301 {
     fmt.Println("$",(income-510301)*.37+(510300-204100)*35+(204100-160725)*.32+(160725-84200)*.24+(84200-39475)*.22+(39475-9700)*.12+(9700*.1),"Is the amount owed and you would be in the 37% tax bracket.")
     }
   }
  if status == "B" && status != "A" && status != "C" && status != "D"{
     fmt.Println("You selected the filing status Married and filing jointly.")
     if income >= 0 && income <= 19400 {
     fmt.Println("$",income*.1,"Is the amount owed and you would be in the 10% tax bracket.")
     }else if income >= 19401 && income <= 78950 {
     fmt.Println("$",(income-19400)*.12+(19400*.1),"Is the amount owed and you would be in the 12% bracket")
     }else if income >= 78951 && income <= 168400 {
     fmt.Println("$",(income-78950)*.22+(78950-19400)*.12+(19400*.1),"Is the amount owed and you would be in the 22% tax bracket.")
     }else if income >= 168401 && income <= 321450 {
     fmt.Println("$",(income-168400)*.24+(168400-78950)*.22+(78950-19400)*.12+(19400*.1),"Is the amount owed and you would be in the 24% tax bracket.")
     }else if income >= 321451 && income <= 408200 {
     fmt.Println("$",(income-321450)*.32+(321450-168400)*.24+(168400-78950)*.22+(78950-19400)*.12+(19400*.1),"is the amount owed and you would be in the 32% tax bracket.")
     }else if income >= 408201 && income <= 612350 {
     fmt.Println("$",(income-408200)*35+(408200-321450)*.32+(321450-168400)*.24+(168400-78950)*.22+(78950-19400)*.12+(19400*.1),"Is the amount owed and you would be in the 35% tax bracket.")
     }else if income >= 612351 {
     fmt.Println("$",(income-612350)*.37+(612350-408200)*.35+(408200-321450)*.32+(321450-168400)*.24+(168400-78950)*.22+(78950-19400)*.12+(19400*.1),"Is the amount owed and you would be in the 37% tax bracket.")
     }
   }
  if status == "C" && status != "A" && status != "B" && status != "D"{
     fmt.Println("You selected the filing status Head of household.")
     if income >= 0 && income <= 13850 {
     fmt.Println("$",income*.1,"Is the amount owed and you would be in the 10% tax bracket.")
     }else if income >= 13851 && income <= 52850 {
     fmt.Println("$",(income-13850)*.12+(13850*.1),"Is the amount owed and you would be in the 12% bracket")
     }else if income >= 52851 && income <= 84200 {
     fmt.Println("$",(income-52850)*.22+(52850-13850)*.12+(13850*.1),"Is the amount owed and you would be in the 22% tax bracket.")
     }else if income >= 84201 && income <= 160700 {
     fmt.Println("$",(income-84200)*.24+(84200-52850)*.22+(52850-13850)*.12+(13850*.1),"Is the amount owed and you would be in the 24% tax bracket.")
     }else if income >= 160701 && income <= 204100 {
     fmt.Println("$",(income-160700)*.32+(160700-84200)*.24+(84200-52850)*.22+(52850-13850)*.12+(13850*.1),"is the amount owed and you would be in the 32% tax bracket.")
     }else if income >= 204101 && income <= 510300 {
     fmt.Println("$",(income-204100)*35+(204100-160700)*.32+(160700-84200)*.24+(84200-52850)*.22+(52850-13850)*.12+(13850*.1),"Is the amount owed and you would be in the 35% tax bracket.")
     }else if income >= 510301 {
     fmt.Println("$",(income-510300)*.37+(510300-204100)*.35+(204100-160700)*.32+(160700-84200)*.24+(84200-52850)*.22+(52850-13850)*.12+(13850*.1),"Is the amount owed and you would be in the 37% tax bracket.")
     }
  }
  if status == "D" && status != "A" && status != "B" && status != "C"{
     fmt.Println("You selected the filing status Married filing separately.")
     if income >= 0 && income <= 9700 {
     fmt.Println("$",income*.1,"Is the amount owed and you would be in the 10% tax bracket.")
     }else if income >= 9701 && income <= 39475 {
     fmt.Println("$",((income-9700)*.12+income*.1),"Is the amount owed and you would be in the 12% bracket")
     }else if income >= 39476 && income <= 84200 {
     fmt.Println("$",((income-39475)*.22)+(39475-9700)*.12+(9700*.1),"Is the amount owed and you would be in the 22% tax bracket.")
     }else if income >= 84201 && income <= 160725 {
     fmt.Println("$",(((income-84200)*.24)+(84200-39475)*.22+((39475-9700)*.12)+((9700)*.1)),"Is the amount owed and you would be in the 24% tax bracket.")
     }else if income >= 160726 && income <= 204100 {
     fmt.Println("$",(income-160725)*.32+(160725-84200)*.24+(84200-39475)*.22+(39475-9700)*.12+(9700*.1),"Is the amount owed and you would be in the 32% tax bracket.")
     }else if income >= 204101 && income <= 306175 {
     fmt.Println("$",(income-204100)*35+(204100-160725)*.32+(160725-84200)*.24+(84200-39475)*.22+(39475-9700)*.12+(9700*.1),"Is the amount owed and you would be in the 35% tax bracket.")
     }else if income >= 306176 {
     fmt.Println("$",(income-306175)*.37+(306175-204100)*.35+(204100-160725)*.32+(160725-84200)*.24+(84200-39475)*.22+(39475-9700)*.12+(9700*.1),"Is the amount owed and you would be in the 37% tax bracket.")
     }
  }
}