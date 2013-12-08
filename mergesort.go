package main

func MergeSort(list []int) {
	mergeSortRange(list, 0, len(list) - 1)
}

func mergeSortRange(list []int, start int, end int) {
	if start < end {
		middle := (start + end) / 2

		mergeSortRange(list, start, middle)
		mergeSortRange(list, middle + 1, end)

		merge(list, start, middle, end)
	}
}

func merge(list []int, start int, middle int, end int) {
	temp := make([]int, 0, end - start + 1)

	index1 := start
	index2 := middle + 1
	index := 0

	for index1 <= middle && index2 <= end {
		if list[index1] < list[index2] {
			temp = append(temp, list[index1])
			index1++
		} else {
			temp = append(temp, list[index2])
			index2++
		}

		index++
	}

	if index1 <= middle {
		temp = append(temp, list[index1:middle + 1]...)
	}

	if index2 <= end {
		temp = append(temp, list[index2:end + 1]...)
	}

	for i := 0; i < len(temp); i++ {
		list[start + i] = temp[i]
	}
}

