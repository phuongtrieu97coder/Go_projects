<img src="https://github.com/phuongtrieu97coder/Readme_Content_Structure/assets/82598726/7e4ee0f6-caa3-45ee-8de2-33b40e93d3c1" alt="golang" width="60px" height="40px">

![image](https://github.com/phuongtrieu97coder/Go_projects/assets/82598726/6fe9f9cc-6bd2-48ff-b049-1df2278870e5)

<a type="button" title="Codecademy_Learn_Go_Course_button" href="https://www.codecademy.com/courses/learn-go-loops-arrays-maps-and-structs/projects/currency-converter" target="_blank" data-CodecademyLearnGoCourseButt="CodecademyLearnGoCourseButt_data"><img src="https://user-images.githubusercontent.com/82598726/175697552-f960b057-9e97-4c3e-a3e2-f2b5f7876de9.png" alt="Codecademy_Learn_Go_Course_button" width="400px" height="300px"></a>

<br><br>

# Currency Converter
<br>

# 1. Introduction:
![Alt text](image.png)
<br>
<br>

# 2. Output:
![Alt text](image-9.png)

![Alt text](image-10.png)
<br>
<br>

# 3. Prompts:

![Alt text](image-1.png)
```go
  currencies := map[string]float32{
    "JPY":130.2,
    "EUR":0.95,
  }
```

![Alt text](image-2.png)
![Alt text](image-3.png)
![Alt text](image-4.png)
![Alt text](image-5.png)
```go
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
```

![Alt text](image-6.png)
![Alt text](image-7.png)
![Alt text](image-8.png)
```go
rate, isValid := currencies[currency]
if isValid {
    fmt.Printf("The amount of %v currency is $%f.\n",currency,rate*dollarAmount)
}else{
    fmt.Println(currency,"currency type was not in the map!")
}
```