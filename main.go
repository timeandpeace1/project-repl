// Esosa Odiase
// 02/13/2020
// a program that will convert your input in kelvins to celcius and farenheit

package main

import "fmt"

func main() {
  fmt.Println("enter distance in rods")
  var first float32
  fmt.Scanln(&first)

  fmt.Println(first-273.15)
  fmt.Println((first-273.15)*(9.0/5)+32)
}