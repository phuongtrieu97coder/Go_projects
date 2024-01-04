<img src="https://github.com/phuongtrieu97coder/Readme_Content_Structure/assets/82598726/7e4ee0f6-caa3-45ee-8de2-33b40e93d3c1" alt="golang" width="60px" height="40px">

![image](https://github.com/phuongtrieu97coder/Go_projects/assets/82598726/6fe9f9cc-6bd2-48ff-b049-1df2278870e5)

<a type="button" title="Codecademy_Learn_Go_Course_button" href="https://www.codecademy.com/courses/learn-go/projects/bank-heist" target="_blank" data-CodecademyLearnGoCourseButt="CodecademyLearnGoCourseButt_data"><img src="https://user-images.githubusercontent.com/82598726/175697552-f960b057-9e97-4c3e-a3e2-f2b5f7876de9.png" alt="Codecademy_Learn_Go_Course_button" width="400px" height="300px"></a>

<br><br>

# Bank Heist
<br>

# 1. Introduction:
![Alt text](image.png)
<br>
<br>

# 2. Output:
> ![Alt text](image-16.png)

> ![Alt text](image-17.png)

> ![Alt text](image-18.png)

> ![Alt text](image-19.png)

> ![Alt text](image-20.png)

<br>
<br>

# 3. Prompts:

![Alt text](image-1.png)
![Alt text](image-2.png)
```go
var isHeistOn bool = true
eludedGuards:= rand.Intn(100)
```

![Alt text](image-3.png)
```go
if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
}
```

![Alt text](image-4.png)
```go
if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  }else{
    isHeistOn = false
     fmt.Println("Plan a better disguise next time?")
}
```


![Alt text](image-5.png)
```go
  fmt.Println(isHeistOn)
```

![Alt text](image-6.png)
```go
  openedVault:=rand.Intn(100)
```

![Alt text](image-7.png)
![Alt text](image-8.png)
```go
if isHeistOn == true && openedVault >=70 {
    fmt.Println("Grab and Go!")
}else if isHeistOn == true && openedVault < 70 {
    isHeistOn = false
    fmt.Println("The vault can't be opened.")
}
```

![Alt text](image-9.png)
```go
  leftSafely:= rand.Intn(5)
```

![Alt text](image-10.png)
![Alt text](image-11.png)
![Alt text](image-12.png)
![Alt text](image-13.png)
![Alt text](image-14.png)
```go
if isHeistOn == true {
    switch leftSafely {
        case 0:
            isHeistOn = false
            fmt.Println("failed heist!!!")
        case 1:
            isHeistOn = false
            fmt.Println("Turns out vault doors don't open from the inside...")
        case 2:
            isHeistOn = false
            fmt.Println("Turns out vault doors don't open from the inside...")
        case 3:
            isHeistOn = false
            fmt.Println("Turns out vault doors don't open from the inside...")
        default:
            fmt.Println("Start the getaway car!")
    }

    if isHeistOn {
      amtStolen:= 10000 + rand.Intn(1000000)
      fmt.Println("$",amtStolen,"not bad!")

    }
}
```

![Alt text](image-15.png)

