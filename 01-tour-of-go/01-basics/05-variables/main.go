package main

import "fmt"

var c, python, java bool
var initalized_c, initalized_python, initalized_java = true, false, "no!"

func main() {
	var i int

	fmt.Println(i, c, python, java)
	fmt.Println(initalized_c, initalized_python, initalized_java)

	short_c, short_python, short_java := true, false, "no!"
	fmt.Println(short_c, short_python, short_java)

	zero_values()
}

func zero_values(){
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}