package main

import (
	"fmt"
)

func main() {
  currencies := map[string]float32{
    "JPY":130.2,
    "EUR":0.95,
  }
  var dollarAmount float32
  var currency string
  fmt.Println("Enter your dollar amount:")
  fmt.Scan(&dollarAmount)

  if dollarAmount == 0 {
    fmt.Println("Invalid amount")
  }else{
     fmt.Println("Your amount is",dollarAmount)
  }

  fmt.Println("Enter your currency type:")

  fmt.Scan(&currency)

  rate, isValid := currencies[currency]
  if isValid {
    fmt.Printf("The amount of %v currency is $%f.\n",currency,rate*dollarAmount)
  }else{
    fmt.Println(currency,"currency type was not in the map!")
  }
}