<img src="https://github.com/phuongtrieu97coder/Readme_Content_Structure/assets/82598726/7e4ee0f6-caa3-45ee-8de2-33b40e93d3c1" alt="golang" width="60px" height="40px">

![image](https://github.com/phuongtrieu97coder/Go_projects/assets/82598726/6fe9f9cc-6bd2-48ff-b049-1df2278870e5)

<a type="button" title="Codecademy_Learn_Go_Course_button" href="https://www.codecademy.com/courses/learn-go-loops-arrays-maps-and-structs/projects/looping-through-a-food-order" target="_blank" data-CodecademyLearnGoCourseButt="CodecademyLearnGoCourseButt_data"><img src="https://user-images.githubusercontent.com/82598726/175697552-f960b057-9e97-4c3e-a3e2-f2b5f7876de9.png" alt="Codecademy_Learn_Go_Course_button" width="400px" height="300px"></a>

<br><br>

# Food Order
<br>

# 1. Introduction:
![Alt text](image.png)
<br>
<br>

# 2. Output:
> $ go run order.go <br>
> The fast food menu has these items: [Burgers Nuggets Fries] 

> What would you like to eat: Burgers

> We have $14 remaining to be paid to the cashier and must hand a $2 bill to the cashier.

> We have $12 remaining to be paid to the cashier and must hand a $2 bill to the cashier.

> We have $10 remaining to be paid to the cashier and must hand a $2 bill to the cashier.

> We have $8 remaining to be paid to the cashier and must hand a $2 bill to the cashier.

> We have $6 remaining to be paid to the cashier and must hand a $2 bill to the cashier.

> We have $4 remaining to be paid to the cashier and must hand a $2 bill to the cashier.

> We have $2 remaining to be paid to the cashier and must hand a $2 bill to the cashier.

> The total for the order is 14
<br>
<br>

# 3. Prompts:

![Alt text](image-1.png)
```go
//In 'main' function assign 'askOrder()' to 'order' variable
var order string = askOrder()


for order!="quit" {
  order=askOrder()
}
```

![Alt text](image-2.png)
![Alt text](image-3.png)
```go
func contains(menu []string,order string) bool {
  for index, value := range menu {
    if order == value {
      return true
    }
  }
  return false
}
//To make the program run correctly, we need to use another for loop through an array
func contains(menu []string,order string) bool {
 
  for i:=0; i<len(menu); i++{
    if order == menu[i] {
      return true
    }
  }
  return false
}

```


![Alt text](image-4.png)
![Alt text](image-5.png)
```go
for order!="quit" {
    if contains(fastfoodMenu, order){
        total+=4
        break
        //without 'break' this loop will run forever
    }else{
        fmt.Println("This item is not on the menu")
        order=askOrder()
    }
}
```
![Alt text](image-6.png)
![Alt text](image-7.png)
![Alt text](image-8.png)
![Alt text](image-9.png)
```go
for amount := total; amount >0; amount-=2 {
    fmt.Printf("We have $%d remaining to be paid to the cashier and must hand a $2 bill to the cashier.\n",amount)
}
```

![Alt text](image-10.png)