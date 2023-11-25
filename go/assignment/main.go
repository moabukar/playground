package main

func main() {
	// crete a slice of int from 0 through 10
	list := []int{7, 3, 4, 5, 6, 7, 8, 9, 10}
	// iterate through the slice and print whether the number is even or odd
	for _, number := range list {
		if number%2 == 0 {
			println(number, "is even")
		} else {
			println(number, "is odd")
		}
	}
}
