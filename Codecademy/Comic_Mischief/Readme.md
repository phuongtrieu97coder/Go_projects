<img src="https://github.com/phuongtrieu97coder/Readme_Content_Structure/assets/82598726/7e4ee0f6-caa3-45ee-8de2-33b40e93d3c1" alt="golang" width="60px" height="40px">

![image](https://github.com/phuongtrieu97coder/Go_projects/assets/82598726/6fe9f9cc-6bd2-48ff-b049-1df2278870e5)

<a type="button" title="Codecademy_Learn_Go_Course_button" href="https://www.codecademy.com/courses/learn-go/projects/comic-mischief" target="_blank" data-CodecademyLearnGoCourseButt="CodecademyLearnGoCourseButt_data"><img src="https://user-images.githubusercontent.com/82598726/175697552-f960b057-9e97-4c3e-a3e2-f2b5f7876de9.png" alt="Codecademy_Learn_Go_Course_button" width="400px" height="300px"></a>

<br><br>

# Comic Mischief
<br>

# 1. Introduction:
![Alt text](image.png)
<br>
<br>

# 2. Output:
> Mr.GoToSleep written by Tracey Hatchet , drawn by Jewel Tampson , published by DizzyBooks Publishing Inc. in 1997 with the number of 14 pages is scoring 6.5

>Epic Vol. 1 written by Ryan N. Shawn , drawn by Phoebe Paperclips , published by DizzyBooks Publishing Inc. in 2013 with the number of 160 pages is scoring 9

<br>
<br>

# 3. Prompts:

![Alt text](image-1.png)
![Alt text](image-2.png)
```go
package main
import "fmt"
func main(){
  
}
```
![Alt text](image-3.png)
```go
  var publisher, writer, artist, title string

```
![Alt text](image-4.png)
![Alt text](image-5.png)
```go
  var year, pageNumber uint

```


![Alt text](image-6.png)
```go
  var grade float32
```


![Alt text](image-7.png)
```go
fmt.Println(title,"written by",writer,", drawn by",artist,", published by",publisher,"in",year,"with the number of",pageNumber,"pages","is scoring",grade)
```

![Alt text](image-8.png)
![Alt text](image-9.png)
![Alt text](image-10.png)
```go
 title = "Mr.GoToSleep"
  writer="Tracey Hatchet"
  artist="Jewel Tampson"
  publisher="DizzyBooks Publishing Inc."
  year=1997
  pageNumber=14
  grade=6.5
```

![Alt text](image-11.png)
![Alt text](image-12.png)
![Alt text](image-13.png)
![Alt text](image-14.png)
![Alt text](image-15.png)
```go
  title = "Epic Vol. 1"
  writer="Ryan N. Shawn"
  artist="Phoebe Paperclips"
  year=2013
  pageNumber=160
  grade=9.0
  fmt.Println(title,"written by",writer,", drawn by",artist,", published by",publisher,"in",year,"with the number of",pageNumber,"pages","is scoring",grade)
```
![Alt text](image-16.png)