// Esosa Odiase
// 02/13/2020
// a program that will convert your input in kelvins to celcius and farenheit

package main

import "fmt"

func main() {
  // you first enter your value in kelvins
  fmt.Println("enter kelvin")
  var first float32
  fmt.Scanln(&first)
  // your value is then converted to celcius then farenheit
  fmt.Println("your value in celcius is", first-273.15)
  fmt.Println("your value in farenheit is", (first-273.15)*(9.0/5)+32)
}