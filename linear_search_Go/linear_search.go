package linear_search

import (
	"errors"
	"fmt"
)

func linear_search(arr []int, target int) (int, error) {
	for i := range arr {
		if arr[i] == target {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func LinearS() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ret, err := linear_search(arr, 11)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Value found at index %v\n", ret)
	}
}
