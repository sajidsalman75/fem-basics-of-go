package main

import "fmt"

// global variable / package variable
// var url = "https://www.google.com"

// func calculateTax(price float32) (float32, float32) {
// 	return price * 0.09, price * 0.15
// }

// func calculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
// 	stateTax = price * 0.09
// 	cityTax = price * 0.15
// 	return
// }

func birthday(age *int) {
	if *age > 140 {
		panic("Too old")
	}
	*age++
}

func main() {
	// function-scoped variables
	// var message string = "Hello from Go"
	// price := 34.4
	// const maxSpeed = 60
	// fmt.Println(message, price, maxSpeed, url)
	// printData()
	// fmt.Println(data.Url)

	// functions
	// stateTax, cityTax := calculateTax(100)
	// fmt.Println(stateTax, cityTax)
	// stateTax1, _ := calculateTax(100)
	// fmt.Println(stateTax1)
	// stateTax2, _ := calculateTaxWithName(100)
	// fmt.Println(stateTax2)

	defer fmt.Println("Good ")
	defer fmt.Println("Bye ")
	// pointer
	var age int = 22
	birthday(&age)
	fmt.Printf("The pointer is %v and the value is %d\n", &age, age)
}
