package main

import "fmt"

func main() {
	simple_for_loop()
	while_loop()
	infinite_loop()
}

func simple_for_loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func while_loop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func infinite_loop() {
	interation := 0

	for {
		interation += 1
		
		if interation == 1000 {
			break
		}

		if interation % 100 == 0 {
			fmt.Println(interation)
		}
	}
}
