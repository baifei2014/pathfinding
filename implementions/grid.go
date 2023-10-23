// @Author jianglonghao
// @Date 2023/10/23
// @Description

package implementions

import "graph/util"

type Graph struct {
	dirs    [4]Location
	width   int
	height  int
	walls   map[Location]bool
	forests map[Location]int
}

func NewGraph(width, height int) Graph {
	graph := Graph{
		dirs:    [4]Location{{1, 0}, {0, -1}, {0, 1}, {-1, 0}},
		width:   width,
		height:  height,
		walls:   map[Location]bool{},
		forests: map[Location]int{},
	}

	return graph
}

func (g Graph) inBounds(id Location) bool {
	return 0 <= id.X && id.X < g.width && 0 <= id.Y && id.Y < g.height
}

func (g Graph) passable(id Location) bool {
	if _, ok := g.walls[id]; ok {
		return false
	}
	return true
}

func (g Graph) neighbors(id Location) []Location {
	var result []Location

	for _, dir := range g.dirs {
		next := Location{
			X: id.X + dir.X,
			Y: id.Y + dir.Y,
		}

		if g.inBounds(next) && g.passable(next) {
			result = append(result, next)
		}

	}
	if (id.X+id.Y)%2 == 0 {
		util.Reverse(result)
	}

	return result
}

func (g Graph) cost(from_node, to_node Location) int {
	_, ok := g.forests[to_node]
	if ok {
		return 5
	}

	return 1
}

func addRect(grid Graph, x1, y1, x2, y2 int) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			grid.walls[Location{x, y}] = true
		}
	}
}

type Location struct {
	X int
	Y int
}

func (l Location) equal(t Location) bool {
	if l.X == t.X && l.Y == t.Y {
		return true
	}

	return false
}
