package main

import (
	"errors"
	"container/list"
	"fmt"
)

type Graph struct {
	adjLists []*list.List
}

func NewGraph(vertices int) (graph *Graph) {
	graph = new(Graph)
	graph.adjLists = make([]*list.List, vertices)

	for i := range graph.adjLists {
		graph.adjLists[i] = list.New()
	}

	return
}

func (graph *Graph) InsertEdge(from int, to int) {
	graph.adjLists[from].PushFront(to)
}

func (graph *Graph) DeleteEdge(from int, to int) {
	for e := graph.adjLists[from].Front(); e != nil; e = e.Next() {
		if e.Value == to {
			graph.adjLists[from].Remove(e)
		}
	}
}

func (graph *Graph) Indegrees() (res []int) {
	res = make([]int, len(graph.adjLists))

	for i := range graph.adjLists {
		for e := graph.adjLists[i].Front(); e != nil; e = e.Next() {
			res[e.Value.(int)]++
		}
	}

	return
}

func (graph *Graph) TopologicalSort() (res []int, err error) {
	temp := graph.Indegrees()
	res = make([]int, 0, len(graph.adjLists))
	queue := list.New()
	visitedCount := 0

	for index, indegree := range temp {
		if indegree == 0 {
			queue.PushBack(index)
		}
	}

	i := 0

	for queue.Front() != nil {
		fmt.Printf("Queue #%d: ", i)
		i++

		for e := queue.Front(); e != nil; e = e.Next() {
			fmt.Printf("%d", e.Value)

			if e.Next() != nil {
				fmt.Printf(", ")
			}
		}
		fmt.Println()

		e := queue.Front()
		vertex := e.Value.(int)

		res = append(res, vertex)
		visitedCount++

		for x := graph.adjLists[vertex].Front(); x != nil; x = x.Next() {
			temp[x.Value.(int)]--

			if temp[x.Value.(int)] == 0 {
				queue.PushBack(x.Value)
			}
		}

		queue.Remove(e)
	}

	if visitedCount != len(graph.adjLists) {
		err = errors.New("The graph is cyclical, topological sort is not possible")
	}

	return
}

func (graph *Graph) ShortestPaths(toVertex int) (pathInfos PathInfos) {
	pathInfos = make(PathInfos, len(graph.adjLists))

	for i := range pathInfos {
		pathInfos[i] = &PathInfo{LastVertex: -1}

		if i == toVertex {
			pathInfos[i].Distance = 0
		} else {
			pathInfos[i].Distance = -1
		}
	}

	queue := list.New()

	queue.PushBack(toVertex)

	for queue.Front() != nil {
		e := queue.Front()
		vertex := e.Value.(int)

		for x := graph.adjLists[vertex].Front(); x != nil; x = x.Next() {
			adjVertex := x.Value.(int)

			if pathInfos[adjVertex].Distance == -1 {
				pathInfos[adjVertex].Distance = pathInfos[vertex].Distance + 1
				pathInfos[adjVertex].LastVertex = vertex
				queue.PushBack(adjVertex)
			}
		}

		queue.Remove(e)
	}

	return
}

type PathInfo struct {
	Distance int
	LastVertex int
}

func (pathInfo *PathInfo) String() string {
	return fmt.Sprintf("&{Distance: %d, LastVertex: %d}", pathInfo.Distance, pathInfo.LastVertex)
}

type PathInfos []*PathInfo

func (pathInfos PathInfos) PrintPath(from int, to int) {
	fmt.Printf("Path from %d to %d: ", from, to)

	pathInfos.printPath(from, to)

	fmt.Println()
}

func (pathInfos PathInfos) printPath(from int, to int) {
	if from == to {
		fmt.Printf("%d", from)
	} else {
		pathInfos.printPath(from, pathInfos[to].LastVertex)

		fmt.Printf(" ‣ %d", to)
	}
}
