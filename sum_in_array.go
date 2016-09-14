package main

import (
	"errors"
	"fmt"
)

func sum_in_array(a []int, s int) ([][]int, error) {
	final_array := make([][]int, 0)
	var i int
	length := len(a)
	if length < 2 {
		return make([][]int, 0), errors.New("Array has less than two values")
	}

	for j := 1; j < length; j++ {
		if a[i]+a[j] == s {
			final_array = append(final_array, []int{a[i], a[j]})
		}

		if j == length-1 {
			i++
			j = i
		}
	}

	return final_array, nil
}

func main() {
	fmt.Println("Hello, playground")

	array := []int{1, 2, 3, 4, 5}
	sum := 7
	output, err := sum_in_array(array, sum)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Output: %#v\n", output)
	}
}
