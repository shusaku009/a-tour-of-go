package main

import (
	"fmt"
	"math"
)

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
	fmt.Println(sqrt(2), sqrt(-4))

}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
