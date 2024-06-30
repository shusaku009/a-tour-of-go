package main

import "fmt"

func main() {
	// for文
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// for文セミコロン省略
	num := 1
	for num < 1000 {
		num += sum
	}
	fmt.Println(num)
}
