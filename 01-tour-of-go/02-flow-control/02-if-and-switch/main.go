package main

import "math/rand"

func main(){
	if_else()
	if_with_a_short_statement()
	switch_statement()
	switch_evaluation_order()
	switch_without_a_condition()
}

func if_else(){
	x := rand.Intn(10)

	if x < 5 {
		println("x is less than 5")
	} else {
		println("x is greater than 5")
	}
}

func if_with_a_short_statement(){
	if x := rand.Intn(10); x < 5 {
		println("x is less than 5")
	} else {
		println("x is greater than 5")
	}
}

func switch_statement(){
	x := rand.Intn(10)

	switch x {
	case 0:
		println("x is 0")
	case 1:
		println("x is 1")
	case 2:
		println("x is 2")
	case 3:
		println("x is 3")
	case 4:
		println("x is 4")
	case 5:
		println("x is 5")
	default:
		println("x is greater than 5")
	}
}

func switch_evaluation_order(){
	x := rand.Intn(10)

	switch x {
		case x % 2:
			println("x is odd")
		default:
			println("x is even")
	}
}

func switch_without_a_condition(){
	x := rand.Intn(10)

	switch {
		case x < 5:
			println("x is less than 5")
		case x == 5:
			println("x is 5")
	}
}