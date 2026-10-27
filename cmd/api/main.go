package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		return "Hello World"
	}
	return "Hello " + name
}

func Sum(numbers [5]int) int {
	sum :=0
	for _, number := range numbers {
		sum+=number
	}

	return sum
}

func main() {
	name := "Sami"
	fmt.Println(Hello(name))
}
