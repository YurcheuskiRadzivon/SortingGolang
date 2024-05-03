package main

import "fmt"

func as(arr []int) {
	arr[0] = 10
}
func main() {
	var arr [2]int = [2]int{1, 2}
	fmt.Println(arr)
	as(arr[:])
	fmt.Println(arr)
	fmt.Println(max(10, max(20, 30)))
}
