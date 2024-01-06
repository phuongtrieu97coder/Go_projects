<img src="https://github.com/phuongtrieu97coder/Readme_Content_Structure/assets/82598726/7e4ee0f6-caa3-45ee-8de2-33b40e93d3c1" alt="golang" width="60px" height="40px">

![image](https://github.com/phuongtrieu97coder/Go_projects/assets/82598726/6fe9f9cc-6bd2-48ff-b049-1df2278870e5)

<a type="button" title="Codecademy_Learn_Go_Course_button" href="https://www.codecademy.com/courses/learn-go-loops-arrays-maps-and-structs/projects/stock-market-golang" target="_blank" data-CodecademyLearnGoCourseButt="CodecademyLearnGoCourseButt_data"><img src="https://user-images.githubusercontent.com/82598726/175697552-f960b057-9e97-4c3e-a3e2-f2b5f7876de9.png" alt="Codecademy_Learn_Go_Course_button" width="400px" height="300px"></a>

<br><br>

# Stock Market
<br>

# 1. Introduction:
![Alt text](image.png)
<br>
<br>

# 2. Output:
> Stock: GOOG . Price: 2313.5

> Stock: AAPL . Price: 157.28

> Stock: FB . Price: 203.77

> Stock: TWTR . Price: 50

> *Update 1:*

> Stock: GOOG . Price: -3117.17

> Stock: AAPL . Price: 157.28

> Stock: FB . Price: 203.77

> Stock: TWTR . Price: 50

> *Update 2:*

> Stock: GOOG . Price: -3117.17

> Stock: AAPL . Price: -3836.227

> Stock: FB . Price: 203.77

> Stock: TWTR . Price: 50


<br>
<br>

# 3. Prompts:
![Alt text](image-1.png)
```go
type Stock struct {
  name string
  price float32
}
```

![Alt text](image-2.png)
![Alt text](image-3.png)
```go
func (stock *Stock) updateStock() {
  change:=randomNumberGen(-10000,10000)
  stock.price=change
}
```

![Alt text](image-4.png)
```go
func displayMarket(market []Stock){
  for i:=0; i<len(market); i++ {
    fmt.Println("Stock:",market[i].name,". Price:",market[i].price)
  }
}
```

![Alt text](image-5.png)
![Alt text](image-6.png)
```go
 stockMarket := []Stock{{"GOOG",2313.50},{"AAPL",157.28},{"FB",203.77},{"TWTR",50.00}}
 // Function calls go here
  displayMarket(stockMarket)
 //update1
  fmt.Println("*Update 1:*")
  stockMarket[0].updateStock()
  displayMarket(stockMarket)
  //update2
   fmt.Println("*Update 2:*")
  stockMarket[1].updateStock()
  displayMarket(stockMarket)
```