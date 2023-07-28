package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func dev(a, b float64) float64 {
	return a / b
}

func display(msg string) {
	fmt.Println(msg)
}

func main() {
	fmt.Println("Hello githib") // Hello githib
	fmt.Println(sum(1, 2))      // 3
	fmt.Println(dev(1, 2))      // 0.5
	display("Hello Gitflow")
}
