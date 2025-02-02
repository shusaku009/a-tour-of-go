package main

import "fmt"

func main() {
	primers := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primers[1:4]
	fmt.Println(s)
}
