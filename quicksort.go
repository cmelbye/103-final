package main

func quickSort(list []int) {
	quickSortRange(list, 0, len(list) - 1)
}

func quickSortRange(list []int, start int, end int) {
	if start < end {
		setPivotToEnd(list, start, end)
		
		pivotIndex := splitList(list, start, end)

		quickSortRange(list, start, pivotIndex - 1)
		quickSortRange(list, pivotIndex + 1, end)
	}
}

// Chooses pivot and moves it to the end of the list segment
// Pivot is the median of the first, the middle, and the last elements.
//
// Precondition: None
// Postcondition: list[end] is the pviot
func setPivotToEnd(list []int, start int, end int) {
	middle := (start + end) / 2

	if list[middle] < list[start] {
		temp := list[start]
		list[start] = list[middle]
		list[middle] = temp
	}

	if list[end] < list[start] {
		temp := list[start]
		list[start] = list[end]
		list[end] = temp
	}

	if list[middle] < list[end] {
		temp := list[middle]
		list[middle] = list[end]
		list[end] = temp
	}
}

func splitList(list []int, start int, end int) int {
	indexL := start
	indexR := end - 1
	pivot := list[end]

	for indexL <= indexR {
		for list[indexL] < pivot {
			indexL++
		}

		for indexL <= indexR && list[indexR] > pivot {
			indexR--
		}

		if indexL <= indexR {
			temp := list[indexL]
			list[indexL] = list[indexR]
			list[indexR] = temp

			indexL++
			indexR--
		}
	}

	temp := list[indexL]
	list[indexL] = list[end]
	list[end] = temp

	return indexL
}
