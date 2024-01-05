<img src="https://github.com/phuongtrieu97coder/Readme_Content_Structure/assets/82598726/7e4ee0f6-caa3-45ee-8de2-33b40e93d3c1" alt="golang" width="60px" height="40px">

![image](https://github.com/phuongtrieu97coder/Go_projects/assets/82598726/6fe9f9cc-6bd2-48ff-b049-1df2278870e5)

<a type="button" title="Codecademy_Learn_Go_Course_button" href="https://www.codecademy.com/courses/learn-go-loops-arrays-maps-and-structs/projects/catnapped" target="_blank" data-CodecademyLearnGoCourseButt="CodecademyLearnGoCourseButt_data"><img src="https://user-images.githubusercontent.com/82598726/175697552-f960b057-9e97-4c3e-a3e2-f2b5f7876de9.png" alt="Codecademy_Learn_Go_Course_button" width="400px" height="300px"></a>

<br><br>

# Catnapped
<br>

# 1. Introduction:
![Alt text](image.png)
![Alt text](image-1.png)
<br>
<br>

# 2. Output:
> Guest List:  [David Guetta Armin Van Buuren Jauz]

> Storage Objects:  [A Toy Chest A Crate A Box]

> Armin Van Buuren hid the cat by putting it in the A Toy Chest
<br>
<br>

# 3. Prompts:
![Alt text](image-2.png)
![Alt text](image-3.png)

```go
 guestList := []string{"David Guetta","Armin Van Buuren","Jauz"}
 storageObjects := []string{"A Toy Chest","A Crate","A Box"}
```

![Alt text](image-4.png)
```go
fmt.Println("Guest List: ",guestList)
fmt.Println("Storage Objects: ",storageObjects)
```

![Alt text](image-5.png)
```go
import (
  "fmt"
  "math/rand"
  "time"
)

//Inside 'main' function add syntaxes below to the start of the function:
 rand.Seed(time.Now().UnixNaNo())
```
> In Go latest verison, rand.Seed() is not required. It is automatically set to the current time.

![Alt text](image-6.png)
![Alt text](image-7.png)
```go
func getRandomElement(array []string) string {
  randomVal := rand.Intn(len(array))
   return array[randomVal]
}
```

![Alt text](image-8.png)
![Alt text](image-9.png)
```go
guestListSlice := guestList[:]
storageObjectsSlice := storageObjects[:]
var Guest string = getRandomElement(guestListSlice)
var Location_Objects string = getRandomElement(storageObjectsSlice)
```
![Alt text](image-10.png)
![Alt text](image-11.png)


