<img src="https://github.com/phuongtrieu97coder/Readme_Content_Structure/assets/82598726/7e4ee0f6-caa3-45ee-8de2-33b40e93d3c1" alt="golang" width="60px" height="40px">

![image](https://github.com/phuongtrieu97coder/Go_projects/assets/82598726/6fe9f9cc-6bd2-48ff-b049-1df2278870e5)

<a type="button" title="Codecademy_Learn_Go_Course_button" href="https://www.codecademy.com/courses/learn-go/projects/learn-go-functions-interstellar-travel" target="_blank" data-CodecademyLearnGoCourseButt="CodecademyLearnGoCourseButt_data"><img src="https://user-images.githubusercontent.com/82598726/175697552-f960b057-9e97-4c3e-a3e2-f2b5f7876de9.png" alt="Codecademy_Learn_Go_Course_button" width="400px" height="300px"></a>

<br><br>

# Interstellar Travel
<br>

# 1. Introduction:
![Alt text](image.png)
<br>
<br>

# 2. Output:
> ![Alt text](image-15.png)

<br>
<br>

# 3. Prompts:

![Alt text](image-1.png)
![Alt text](image-2.png)
```go
func fuelGauge(fuel int){
  fmt.Println("You have",fuel," fuel left")
}
```

![Alt text](image-3.png)
```go
func calculateFuel(planet string) int{
  var fuel int
}
```

![Alt text](image-4.png)
![Alt text](image-5.png)
```go
func calculateFuel(planet string) int{
  var fuel int
  switch planet {
    case "Venus":
      fuel = 300000
    case "Mercury":
      fuel = 500000
    case "Mars":
      fuel = 700000
    default:
      fuel = 0
  }
  return fuel
}
```

![Alt text](image-6.png)
![Alt text](image-7.png)
![Alt text](image-8.png)
```go
// Create the function greetPlanet() here
func greetPlanet(planet string){
  fmt.Println("Welcome to planet",planet)
}

// Create the function cantFly() here
func cantFly(){
  fmt.Println("We do not have the available fuel to fly there.")
}
```

![Alt text](image-9.png)
![Alt text](image-10.png)
![Alt text](image-11.png)
```go
// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)

  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining-=fuelCost
  }else{
    cantFly()
  }
  return fuelRemaining
}
```

![Alt text](image-12.png)
![Alt text](image-13.png)
```go
func main() {
  // Test your functions!
  
  // Create `planetChoice` and `fuel`
  var fuel int = 1000000
  planetChoice := "Venus"

  fuel = flyToPlanet(planetChoice,fuel)
  fuelGauge(fuel)
  // And then liftoff!
  
}
```

![Alt text](image-14.png)

