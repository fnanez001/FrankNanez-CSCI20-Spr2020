// Frank Nanez
// 3-3-20
// products and pricing with 2 arrays

package main

import "fmt"

func main() {
    //Create a string array of at least 5 values.  It should hold 5 products to buy
    var products [5]string
    products[0] = "milk"
    products[1] = "gel"
    products[2] = "flour"
    products[3] = "toothbrush"
    products[4] = "tortilla"
  
    //Create a float array of at least 5 values.  It should hold a price for each of the products
    var price [5]float32
    price[0] = 3.99
    price[1] = 3.99
    price[2] = 1.79
    price[3] = .99
    price[4] = 2.99
    //Iterate through the array and output the product and it's price (similar to a menu)
    for i := 0; i <= 4; i++ {
      fmt.Println(products[i], price[i])
    }
}