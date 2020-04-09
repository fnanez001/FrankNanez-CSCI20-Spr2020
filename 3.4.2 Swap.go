// Frank Nanez
// 4-8-20
// takes user numbers and swaps them. prints out before and after of numbers given

package main

import "fmt"

func swap (input1 int, input2 int){
 fmt.Println("Initial input:", input1,input2)
 fmt.Println("Swapped numbers:",input2,input1)
}

func main () {
 var num1 int
 var num2 int

 fmt.Println("Please enter 2 numbers and they will be swapped for you. The first number is?")
 fmt.Scanln(&num1)
 fmt.Println("The second number is?")
 fmt.Scanln(&num2)

 swap(num1,num2)
}