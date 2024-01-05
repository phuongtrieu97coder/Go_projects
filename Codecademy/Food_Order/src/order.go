package main

import (
  "fmt"
)

func askOrder() string {
  var input string
  fmt.Print("What would you like to eat: ")
  // Get the input from the user
  fmt.Scan(&input)
  return input
}

// WRITE CONTAINS FUNCTION BELOW
func contains(menu []string,order string) bool {
  /*
  for index, value := range menu {
    if order == value {
      return true
    }
  }
  */
  for i:=0; i<len(menu); i++{
    if order == menu[i] {
      return true
    }
  }
  return false
}


func main() {
  fastfoodMenu := []string{"Burgers", "Nuggets", "Fries"}
  fmt.Println("The fast food menu has these items:", fastfoodMenu)
 
  var total int = 10
  var order string = askOrder()
  
  // WRITE INDEFINITE LOOP ASKING FOR ORDERS BELOW
  for order != "quit" {
    if contains(fastfoodMenu, order){
      total+=4
      break
    }else{
      fmt.Println("This item is not on the menu")
      order = askOrder()
    }
  }

 

  // WRITE DEFINITE LOOP COUNTING $2 BILLS BELOW
  for amount := total; amount >0; amount-=2 {
      fmt.Printf("We have $%d remaining to be paid to the cashier and must hand a $2 bill to the cashier.\n",amount)
  }


  fmt.Println("The total for the order is", total)
}
