// append() 和 copy() 函数
// 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
// 下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。
package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量 cap按最大值的2倍扩充*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*5)
	printSlice(numbers1)
	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
	var i int
	for i = 0; i < len(numbers); i++ {
		fmt.Println("###")
		numbers1 = append(numbers1, numbers[i])
		printSlice(numbers1)
		numbers1 = append(numbers1, numbers[i])
		printSlice(numbers1)
		numbers1 = append(numbers1, numbers[i])
		printSlice(numbers1)
		numbers1 = append(numbers1, numbers[i])
		printSlice(numbers1)
	}
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
