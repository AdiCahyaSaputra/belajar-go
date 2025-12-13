package main

import "fmt"

func sayHelloTo(name string) {
	fmt.Println("Hello", name)
}

func getHello(name string) string {
	return "Hello " + name
}

// GELO BISA GINI ANJAY
func getFullName() (string, string) {
	return "Adi", "Cahya"
}

// Named Return Value tapi wajib sebut tipe data
func getCompleteName() (firstName, lastName string) {
	firstName = "Adi"
	lastName = "Cahya"

	return firstName, lastName
}

// Rest Parameter (JS) -> Variadic Function (Go)
func sumAll(numbers ...int) int {
	total := 0

	// numbers disini bakal jadi slice
	for _, number := range numbers {
		total += number
	}

	return total
}

// Callback function
func filterNumber(numbers []int, callback func(int) bool) []int {
	var result []int

	for _, number := range numbers {
		if callback(number) {
			result = append(result, number)
		}
	}

	return result
}

// Recursive Function
func factorialLoop(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
}

func main() {
	sayHelloTo("Adi")

	result := getHello("Adi")
	fmt.Println(result)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	firstName, lastName = getCompleteName()
	fmt.Println(firstName, lastName)

	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	// ATAU
	numbers := []int{10, 10, 10, 10, 10}

	// ... disini untuk mengubah slice menjadi variadic / slice nya kayak di destructure gitu lho
	total = sumAll(numbers...)
	fmt.Println(total)

	// Callback function
	// btw disini juga ada anonymous function
	resultSliceInt := filterNumber(numbers, func(number int) bool {
		return number > 10
	})

	fmt.Println(resultSliceInt)

	factorialValue := factorialLoop(5)
	fmt.Println(factorialValue)
}
