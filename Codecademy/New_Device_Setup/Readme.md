<img src="https://github.com/phuongtrieu97coder/Readme_Content_Structure/assets/82598726/7e4ee0f6-caa3-45ee-8de2-33b40e93d3c1" alt="golang" width="60px" height="40px">

![image](https://github.com/phuongtrieu97coder/Go_projects/assets/82598726/6fe9f9cc-6bd2-48ff-b049-1df2278870e5)

<a type="button" title="Codecademy_Learn_Go_Course_button" href="https://www.codecademy.com/courses/learn-go/projects/new-device-setup-with-go-fmt" target="_blank" data-CodecademyLearnGoCourseButt="CodecademyLearnGoCourseButt_data"><img src="https://user-images.githubusercontent.com/82598726/175697552-f960b057-9e97-4c3e-a3e2-f2b5f7876de9.png" alt="Codecademy_Learn_Go_Course_button" width="400px" height="300px"></a>

<br><br>

# New Device Setup
<br>

# 1. Introduction:
![Alt text](image.png)
<br>
<br>

# 2. Output:
> ![Alt text](image-7.png)


<br>
<br>

# 3. Prompts:

![Alt text](image-1.png)
![Alt text](image-2.png)
```go
var name string
  fmt.Println("What's your name?")
  fmt.Scan(&name)
  fmt.Printf("Hello %s!!!\n",name)

```

![Alt text](image-3.png)
![Alt text](image-4.png)
```go
 var age int
  fmt.Println("What's your age?")
  fmt.Scan(&age)
  fmt.Printf("Age: %d.\n",age)
```

![Alt text](image-5.png)
![Alt text](image-6.png)
```go
newMessage:=fmt.Sprintf("Name: %s. Age: %d",name,age)
fmt.Println(newMessage)

```