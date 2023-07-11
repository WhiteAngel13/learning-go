package main

import "fmt"

func add_with_each_values_typed(x int, y int) int {
	return x + y
}

func add_with_all_values_typed(x, y int) int {
	return x + y
}

func multi_value_returned_swap(x, y string) (string, string) {
	return y, x
}

func named_return_values_split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add_with_each_values_typed(42, 13))
	fmt.Println(add_with_all_values_typed(42, 13))
	a, b := multi_value_returned_swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(named_return_values_split(84))
}
