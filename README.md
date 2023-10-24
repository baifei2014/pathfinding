
pathfinding library about dijkstra and a_star algorithm

## Examples

The following are demonstration examples of two algorithms

### Dijkstra

#### first step

- generate graph
- define start, goal
- define cost of cost so far and shortest path by came from

```go
graph := implementions.MakeDiagram4()

start := implementions.Location{X: 1, Y: 4}
goal := implementions.Location{X: 8, Y: 5}

came_from := map[implementions.Location]implementions.Location{}
cost_so_far := map[implementions.Location]int{}
```

#### second step

- do search
- construct path

```go
implementions.DijkstraSearch(graph, start, goal, came_from, cost_so_far)

path := implementions.ReconstructPath(start, goal, came_from)
```

#### third step

- draw path on graph

```go
implementions.DrawGrid(graph, path)
```

### A_Star|A*

#### first step

- generate graph
- define start, goal
- define cost of cost so far and shortest path by came from

```go
graph := implementions.MakeDiagram4()

start := implementions.Location{X: 1, Y: 4}
goal := implementions.Location{X: 8, Y: 5}

came_from := map[implementions.Location]implementions.Location{}
cost_so_far := map[implementions.Location]int{}
```

#### second step

- do search
- construct path

```go
implementions.AStarSearch(graph, start, goal, came_from, cost_so_far)

path := implementions.ReconstructPath(start, goal, came_from)
```

#### third step

- draw path on graph

```go
implementions.DrawGrid(graph, path)
```

