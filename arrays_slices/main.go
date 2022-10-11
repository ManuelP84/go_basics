package main

import (
	"fmt"
)

func main() {

	// Array
	fmt.Println()
	fmt.Println("::: Array :: cap()/len() :::")
	fmt.Println()
	array1 := [7]int{1, 3, 5, 7, 9, 11, 13}
	array2 := array1[2:4]
	array2 = append(array2, 6)
	fmt.Println("Array 1")
	fmt.Println(array1, len(array1), cap(array1))
	fmt.Println("Array 2")
	fmt.Println(array2, len(array2), cap(array2))

	// Slice
	fmt.Println()
	fmt.Println("::: Slice :: cap()/len() :::")
	fmt.Println()
	slice := []int{2, 4, 6, 8, 10, 12, 14}
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 16)
	fmt.Println(slice, len(slice), cap(slice))

	// Slice methods
	fmt.Println()
	fmt.Println("::: Slice methods :::")
	fmt.Println()
	fmt.Println("slice[2]: ", slice[2])
	fmt.Println("slice[:3]: ", slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	fmt.Println()
	fmt.Println(":::Append:::")
	fmt.Println()
	slice = append(slice, 16, 18, 20)
	slice_copy := slice
	slice_copy = append(slice_copy, 1000)
	fmt.Println("slice: ", slice)
	fmt.Println("slice_copy", slice_copy)
	newSlice := []int{22, 24, 26, 28, 30}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// Test
	fmt.Println()
	fmt.Println(":::Test:::")
	fmt.Println()
	test_slice1 := []int{1, 3, 5}
	test_slice2 := test_slice1[:]
	fmt.Println(test_slice1)
	fmt.Println(test_slice2)
	test_slice1[0] = 10
	fmt.Println(test_slice1)
	fmt.Println(test_slice2)
	test_slice1 = append(test_slice1, 50)
	fmt.Println(test_slice1)
	fmt.Println(test_slice2)

	copy_test_slice := make([]int, 4, 5)
	copy(copy_test_slice, test_slice1)
	test_slice1[0] = 100
	fmt.Println(test_slice1)
	fmt.Println(copy_test_slice)

}
