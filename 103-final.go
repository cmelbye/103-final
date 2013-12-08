package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {
	sorts()
	graphs()
}

func sorts() {
	section("Sorting Algorithms")

	originalList := []int{31, 10, 8, 7, 6, 5, 9, 1, 13, 21}

	fmt.Println("Before:", originalList)

	mergeSortList := copyList(originalList)
	MergeSort(mergeSortList)
	fmt.Println("After merge sort:", mergeSortList)

	quickSortList := copyList(originalList)
	QuickSort(quickSortList)
	fmt.Println("After quick sort:", quickSortList)

	fmt.Println()
}

func graphs() {
	section("Graphs and Graph Algorithms")

	graph := NewGraph(7)

	graph.InsertEdge(0, 3)
	graph.InsertEdge(0, 4)
	graph.InsertEdge(1, 0)
	graph.InsertEdge(4, 2)
	graph.InsertEdge(4, 6)
	graph.InsertEdge(5, 4)
	graph.InsertEdge(5, 6)
	graph.InsertEdge(6, 2)
	graph.InsertEdge(6, 3)

	fmt.Println("Indegrees:", graph.Indegrees())
	fmt.Println()

	sort, err := graph.TopologicalSort()

	if err == nil {
		fmt.Println("Topological sort:", sort)
	} else {
		fmt.Println("Error while sorting:", err.Error())
	}

	fmt.Println()

	// Add a few more edges which make the graph cyclical.
	graph.InsertEdge(2, 1)
	graph.InsertEdge(3, 4)

	pathInfos := graph.ShortestPaths(3)

	for _, pathInfo := range pathInfos {
		fmt.Println(pathInfo)
	}

	pathInfos.PrintPath(3, 0)

	fmt.Println()
}


func section(header string) {
	fmt.Printf("%s\n", strings.ToUpper(header))

	for i := 0; i < utf8.RuneCountInString(header); i++ {
		fmt.Print("-")
	}

	fmt.Println()
	fmt.Println()
}

func copyList(list []int) (newList []int) {
	newList = make([]int, len(list))
	copy(newList, list)
	return
}
