// https://www.hackerrank.com/challenges/insertionsort1/problem

package main

import "fmt"

func printArray(arr []int32) {
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}

	fmt.Println()
}

func insertionSort1(n int32, arr []int32) {
	for i := n - 1; i > 0; i-- {
		var temp = arr[i]

		for j := i - 1; j >= 0; j-- {
			if temp < arr[j] {
				arr[j+1] = arr[j]
				printArray(arr)

				if j == 0 {
					arr[0] = temp
					printArray(arr)
				}
			} else {
				if j < i-1 {
					arr[j+1] = temp
					printArray(arr)
				}
				break
			}
		}
	}
}

func main() {
	//insertionSort1(5, []int32{2, 4, 6, 8, 3})
	insertionSort1(10, []int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 1})
}
