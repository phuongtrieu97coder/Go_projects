package main

import (
  "fmt"
  "math/rand"
 // "time"
)

func getRandomElement(array []string) string {
  randomVal := rand.Intn(len(array))
  return array[randomVal]
}

func main() {
  //rand.Seed(time.Now().UnixNaNo())

  guestList := []string{"David Guetta","Armin Van Buuren","Jauz"}
  storageObjects := []string{"A Toy Chest","A Crate","A Box"}
  fmt.Println("Guest List: ",guestList)
  fmt.Println("Storage Objects: ",storageObjects)

  guestListSlice := guestList[:]
  storageObjectsSlice := storageObjects[:]
  var Guest string = getRandomElement(guestListSlice)
  var Location_Objects string = getRandomElement(storageObjectsSlice)
  fmt.Println(Guest,"hid the cat by putting it in the",Location_Objects)
}