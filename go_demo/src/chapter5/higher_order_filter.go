package main

import "fmt"

func IntFilter(slice []int, predicate func(int) bool) []int {
	filtered := make([]int, 0, len(slice))
	for i := 0; i < len(slice); i++ {
		if predicate(slice[i]) {
			filtered = append(filtered, slice[i])
		}
	}
	return filtered
}

func Filter(limit int, predicate func(int) bool, appender func(int)) {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			appender(i)
		}
	}
}

func main() {

	readings := []int{4, -3, 2, -7, 8, 19, -11, 7, 18, -6}
	even := IntFilter(readings, func(i int) bool { return i%2 == 0 })
	fmt.Println(even)
	// 通用调用
	// filter int
	readings1 := []int{4, -3, 2, -7, 8, 19, -11, 7, 18, -6}
	even1 := make([]int, 0, len(readings1))
	Filter(len(readings1), func(i int) bool { return readings1[i]%2 == 0 },
		func(i int) { even1 = append(even1, readings1[i]) })
	fmt.Println(even1)

	// filter string
	parts := []string{"X15", "T14", "X23", "A41", "L19", "x57", "A63"}
	var Xparts []string
	Filter(len(parts), func(i int) bool { return parts[i][0] == 'X' },
		func(i int) { Xparts = append(Xparts, parts[i]) })
	fmt.Println(Xparts)
	// int64
	var product int64 = 1
	productSlice := make([]int, 0, len(readings1))
	Filter(26, func(i int) bool { return i%2 != 0 },
		func(i int) {
			product *= int64(i)
			productSlice = append(productSlice, i)
		})
	fmt.Println(productSlice)
	fmt.Println(product)
}
