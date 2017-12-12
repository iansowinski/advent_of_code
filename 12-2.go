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
	groupsCounter := 0
	allKeys := make([]int, 0)
	for key := range graph {
		allKeys = append(allKeys, key)
	}
	for len(allKeys) != 0 {
		currentLinked := getConnectedNodesList(allKeys[0], graph, connectedList)
		allKeys = difference(allKeys, currentLinked)
		groupsCounter += 1
	}
	fmt.Println(groupsCounter)
}

func difference(a, b []int) []int {
	mb := map[int]bool{}
	for _, x := range b {
		mb[x] = true
	}
	ab := []int{}
	for _, x := range a {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}
	return ab
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
