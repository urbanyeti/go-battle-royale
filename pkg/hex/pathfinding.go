package hex

import (
	"container/heap"
	"fmt"
)

type TileItem struct {
	*Tile
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type TileFrontier []*TileItem

func (tf TileFrontier) Len() int { return len(tf) }

func (tf TileFrontier) Less(i, j int) bool {
	// We want Pop to give us the lowest priority
	return tf[i].priority < tf[j].priority
}

func (tf TileFrontier) Swap(i, j int) {
	tf[i], tf[j] = tf[j], tf[i]
	tf[i].index = i
	tf[j].index = j
}

func (tf *TileFrontier) Push(x interface{}) {
	n := len(*tf)
	item := &TileItem{x.(*Tile), 0, n}
	*tf = append(*tf, item)
}

func (tf *TileFrontier) PushPriority(t *Tile, priority int) {
	n := len(*tf)
	item := &TileItem{t, priority, n}
	*tf = append(*tf, item)
	heap.Fix(tf, item.index)
}

func (tf *TileFrontier) Pop() interface{} {
	old := *tf
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*tf = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (tf *TileFrontier) update(item *TileItem, value *Tile, priority int) {
	item.Tile = value
	item.priority = priority
	heap.Fix(tf, item.index)
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func cubeToAxial(c Cube) Coordinate {
	return Coordinate{c.X, c.Z}
}

func axialToCube(c Coordinate) Cube {
	return Cube{c.Q, (c.Q * -1) - c.R, c.R}
}

func cubeDistance(a Cube, b Cube) int32 {
	return (abs(a.X-b.Y) + abs(a.Y-b.Y) + abs(a.Z-b.Z)) / 2
}

func Distance(a Coordinate, b Coordinate) int32 {
	return cubeDistance(axialToCube(a), axialToCube(b))
}

func TravelCost(a Coordinate, b *Tile) int {
	return int(Distance(a, b.Coordinate) * b.Cost)
}

func FindPath(h HexService, start Coordinate, goal Coordinate) {
	frontier := TileFrontier{}
	frontier.Push(h.Battlefield[start])
	cameFrom := make(map[Coordinate]Coordinate)
	costSoFar := make(map[Coordinate]int)

	costSoFar[start] = 0

	for frontier.Len() != 0 {
		current := heap.Pop(&frontier).(*TileItem)

		if current.Coordinate == goal {
			break
		}

		for _, next := range h.GetNeighbors(current.Tile) {
			if next == nil {
				continue
			}
			newCost := costSoFar[current.Coordinate] + int(next.Cost)

			if val, ok := costSoFar[next.Coordinate]; !ok || newCost < val {
				costSoFar[next.Coordinate] = newCost
				priority := newCost + TravelCost(current.Coordinate, next)
				frontier.PushPriority(next, priority)
				cameFrom[next.Coordinate] = current.Coordinate
			}
		}
	}

	current := goal
	var path []Coordinate
	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	path = append(path, start)

	for i := len(path) - 1; i >= 0; i-- {
		fmt.Println(path[i])
	}

}
