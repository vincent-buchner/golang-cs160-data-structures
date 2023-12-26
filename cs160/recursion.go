package main 

import "fmt"

// function Numbers 
// parameters is integer startingNumber 
// returns a value int 
// returns the sum of numbers from 0 to 
// the given startingNumber 

func Numbers(startingNumber int) int {
  if startingNumber == 0 {
    return startingNumber
  }
  return startingNumber + Numbers(startingNumber - 1)
}

func Factorial(number int) int {
  if number <= 1 {
    return number
  }
  return number * Factorial(number -1)
}

func Fibonacci(number int) int {
  if number <= 1 {
    return number
  }
  return Fibonacci(number -1) + Fibonacci(number - 2)
}

func main(){

  resultNumbers := Numbers(5)
  fmt.Println("Numbers Sum is : ",resultNumbers)
  
  factorialNumber := Factorial(100)
  fmt.Println("Factorial for 100 is : ",factorialNumber)

  fibonacci := Fibonacci(10)
  fmt.Println("Fibonacci Number for 10th position is : ",fibonacci)

}
