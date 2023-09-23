package main

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{12, 11, 13}
	arr_result := []int{11, 12, 13}

	insertionSort(arr)

	if len(arr) != len(arr_result) {
		t.Errorf("arr [%v] and arr_result[%v] are not equal", arr, arr_result)
	}

        for i := range arr {
		if arr[i] != arr_result[i] {
			t.Errorf("arr [%v] and arr_result[%v] are not equal", arr, arr_result)
			break
		}
        }
}


func TestInsertionSortWithEmptySlice(t *testing.T) {
	arr := []int{}
	arr_result := []int{}

	insertionSort(arr)

	if len(arr) != len(arr_result) {
		t.Errorf("arr [%v] and arr_result[%v] are not equal", arr, arr_result)
	}

        for i := range arr {
		if arr[i] != arr_result[i] {
			t.Errorf("arr [%v] and arr_result[%v] are not equal", arr, arr_result)
			break
		}
        }
}


func TestInsertionSortVeryLong(t *testing.T) {
	arr := []int{116, 111, 106, 121, 131, 126, 1, 2, 3, 4, 5, 1, 2, 3, 4, 101, 5}
	arr_result := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 101, 106, 111, 116, 121, 126, 131}

	insertionSort(arr)

	if len(arr) != len(arr_result) {
		t.Errorf("arr [%v] and arr_result[%v] are not equal", arr, arr_result)
	}

        for i := range arr {
		if arr[i] != arr_result[i] {
			t.Errorf("arr [%v] and arr_result[%v] are not equal", arr, arr_result)
			break
		}
        }
}
