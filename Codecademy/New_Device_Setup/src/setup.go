package main

import (
  "fmt"
)

func main() {
  var name string
  fmt.Println("What's your name?")
  fmt.Scan(&name)
  fmt.Printf("Hello %s!!!\n",name)

  var age int
  fmt.Println("What's your age?")
  fmt.Scan(&age)
  fmt.Printf("Age: %d.\n",age)

  newMessage:=fmt.Sprintf("Name: %s. Age: %d",name,age)
  fmt.Println(newMessage)
}