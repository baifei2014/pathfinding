// @Author jianglonghao
// @Date 2023/10/23
// @Description

package implementions

import (
	"fmt"
	"graph/util"
)

func ReconstructPath(start, goal Location, came_from map[Location]Location) []Location {
	var path []Location

	current := goal

	for !current.equal(start) {
		path = append(path, current)
		current = came_from[current]
	}

	path = append(path, start)
	util.Reverse(path)
	return path
}

func MakeDiagram4() Graph {
	grid := NewGraph(10, 10)
	addRect(grid, 1, 7, 4, 9)

	grid.forests = map[Location]int{
		{3, 4}: 1, {3, 5}: 1, {4, 1}: 1, {4, 2}: 1,
		{4, 3}: 1, {4, 4}: 1, {4, 5}: 1, {4, 6}: 1,
		{4, 7}: 1, {4, 8}: 1, {5, 1}: 1, {5, 2}: 1,
		{5, 3}: 1, {5, 4}: 1, {5, 5}: 1, {5, 6}: 1,
		{5, 7}: 1, {5, 8}: 1, {6, 2}: 1, {6, 3}: 1,
		{6, 4}: 1, {6, 5}: 1, {6, 6}: 1, {6, 7}: 1,
		{7, 3}: 1, {7, 4}: 1, {7, 5}: 1,
	}

	return grid
}

func print(s string) {
	fmt.Printf("%-3s", s)
}

func DrawGrid(graph Graph, path []Location) {
	newpath := map[Location]bool{}
	for _, location := range path {
		newpath[location] = true
	}

	for y := 0; y != graph.height; y++ {
		for x := 0; x != graph.width; x++ {
			id := Location{x, y}
			if _, ok := graph.walls[id]; ok {
				print("#")
			} else if _, ok := newpath[id]; ok {
				print("@")
			} else {
				print(".")
			}
		}
		fmt.Println()
	}
}
