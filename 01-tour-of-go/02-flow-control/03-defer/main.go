package main

func main(){
	standard_defer()
	defer_in_a_loop()
}

func standard_defer(){
	defer println("deferred")
	println("not deferred")
}

func defer_in_a_loop(){
	for i := 0; i < 10; i++ {
		defer println(i)
	}

	println("done")
}
