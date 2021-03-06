package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	id           int
	children     []node
	childrenTemp []int
}

func main() {
	input := `` //Here put the input
	graph := populate(input)
	var root int
	root = 0
	connectedList := make([]int, root)
	fmt.Println(len(getConnectedNodesList(0, graph, connectedList)))
}

func getConnectedNodesList(root int, graph map[int][]int, connectedList []int) []int {
	for _, item := range graph[root] {
		if !itemInSlice(item, connectedList) {
			connectedList = append(connectedList, item)
			connectedList = getConnectedNodesList(item, graph, connectedList)
		}
	}
	return connectedList
}

func itemInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func populate(input string) map[int][]int {
	nodeMapTemp := strings.Split(input, "\n")
	nodeMap := make(map[int][]int, len(nodeMapTemp))
	for index, item := range nodeMapTemp {
		firstSplit := strings.Split(item, " <-> ")
		childrenIds := strings.Split(firstSplit[1], ", ")
		childrenList := make([]int, len(childrenIds))
		for childIndex, childNodeId := range childrenIds {
			id, _ := strconv.Atoi(childNodeId)
			childrenList[childIndex] = id
		}
		nodeMap[index] = childrenList
	}
	return nodeMap
}
