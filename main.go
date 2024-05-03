package main

import (
	"fmt"
	"math/rand"
)

func Swap(slice interface{}, i, j int) {
	s := slice.([]int)
	s[i], s[j] = s[j], s[i]
}
func scan(x int) int {
	fmt.Scan(&x)
	return x
}
func zap(arr []int, x int) []int {
	for i := 0; i < x; i++ {
		arr = append(arr, rand.Int()%1000000)
	}
	return arr
}
func copysl(stru data, arr1 []int) []int {
	arr2 := make([]int, len(arr1))
	copy(arr2, arr1)
	return arr2

}

func siftDown(arr []int, size int, x int, stru *data) {
	i := x
	l, r := x*2+1, x*2+2

	stru.colsr++
	if l < size && arr[l] > arr[i] {
		i = l
	}
	if r < size && arr[r] > arr[i] {
		i = r
	}
	if i != x {
		stru.colper++
		Swap(arr, x, i)
		siftDown(arr, size, i, stru)
	}

}

func sortvst(arr []int) {

	fmt.Println("Сортировка методом включения")
	var sort data
	sort.ardo = copysl(sort, arr)
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			sort.colsr++
			if arr[j] < arr[j-1] {
				sort.colper++
				Swap(arr, j-1, j)
			}

			j--

		}
	}
	sort.arposle = arr
	output(sort)

}
func sortbyselect(arr []int) {
	fmt.Println("Сортировка методом выбора")
	var sort data
	sort.ardo = copysl(sort, arr)
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			sort.colsr++
			if arr[min] > arr[j] {
				sort.colper++
				min = j
			}
		}
		sort.colper++
		Swap(arr, i, min)

	}
	sort.arposle = arr
	output(sort)

}

func heapsort(arr []int) {
	fmt.Println("Пирамидальная сортировка")
	var sort data
	sort.ardo = copysl(sort, arr)
	for i := len(arr)/2 - 1; i >= 0; i-- {
		siftDown(arr, len(arr), i, &sort)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		Swap(arr, 0, i)
		siftDown(arr, i, 0, &sort)
	}
	sort.arposle = arr
	output(sort)
}
func shakersort(arr []int) {
	fmt.Println("Шейкерная сортировка")
	var sort data
	sort.ardo = copysl(sort, arr)
	l, r := 0, len(arr)-1
	flag := 1
	for l < r && flag > 0 {
		flag = 0
		for i := l; i < r; i++ {
			sort.colsr++
			if arr[i] > arr[i+1] {
				sort.colper++
				Swap(arr, i, i+1)
				flag = 1
			}

		}
		r--
		for j := r; j > l; j-- {
			sort.colsr++
			if arr[j] > arr[j+1] {
				sort.colper++
				Swap(arr, j, j+1)
				flag = 1
			}
		}
		l++
		if flag == 0 {
			break
		}

	}
	sort.arposle = arr
	output(sort)

}

type data struct {
	ardo    []int
	arposle []int
	colper  int
	colsr   int
}

func output(stru data) {

	fmt.Println("До:")
	fmt.Println(stru.ardo)
	fmt.Println("После:")
	fmt.Println(stru.arposle)
	fmt.Printf("%d - количество перестановок, %d - количество сравнений \n", stru.colper, stru.colsr)
}
func copys(arr1 []int) []int {
	arr2 := make([]int, len(arr1))
	copy(arr2, arr1)
	return arr2
}
func main() {

	//rand.Seed(time.Now().UnixNano())
	n := 0
	n = scan(n)
	var arr1 []int
	arr1 = zap(arr1, n)
	arr2 := copys(arr1)
	arr3 := copys(arr1)
	arr4 := copys(arr1)
	sortvst(arr1)
	sortbyselect(arr2)
	heapsort(arr3)
	shakersort(arr4)

}
