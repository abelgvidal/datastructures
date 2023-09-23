package main

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		currentElement := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > currentElement {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = currentElement
	}
}
