// @Author jianglonghao
// @Date 2023/10/23
// @Description

package implementions

import (
	"github.com/baifei2014/jqueue/priority"
	"math"
)

func heuristic(a, b Location) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func AStarSearch(graph Graph, start, goal Location, came_from map[Location]Location, cost_so_far map[Location]int) {
	frontier := priority.New(priority.Lesser)
	frontier.Put(start, 0)

	came_from[start] = start
	cost_so_far[start] = 0

	for !frontier.Empty() {
		current, ok := frontier.Get().(Location)
		if !ok {
			continue
		}

		if current.equal(goal) {
			break
		}

		for _, next := range graph.neighbors(current) {
			new_cost := cost_so_far[current] + graph.cost(current, next)
			if _, ok := cost_so_far[next]; !ok || new_cost < cost_so_far[next] {
				cost_so_far[next] = new_cost
				priority := new_cost + heuristic(next, goal)
				frontier.Put(next, priority)
				came_from[next] = current
			}
		}
	}
}
