package hex

import (
	"fmt"
	"os"
)

type Direction int

const (
	E Direction = iota
	NE
	NW
	W
	SW
	SE
)

type Directions []*Coordinate

func LoadDirections() Directions {
	c := Directions{
		{Q: 1, R: 0},
		{Q: 1, R: -1},
		{Q: 0, R: -1},
		{Q: -1, R: 0},
		{Q: -1, R: 1},
		{Q: 0, R: 1},
	}
	return c
}

func (h HexService) GetNeighbor(t *Tile, d Direction) *Tile {
	dir := h.Directions[d]
	c := t.Coordinate
	return h.Battlefield[Coordinate{Q: c.Q + dir.Q, R: c.R + dir.R}]
}

func (h HexService) GetNeighbors(t *Tile) []*Tile {
	n := []*Tile{}
	for i := 0; i < len(h.Directions); i++ {
		n = append(n, h.GetNeighbor(t, Direction(i)))
	}
	return n
}

type HexService struct {
	Directions  Directions
	Battlefield Battlefield
}

func NewHexService() HexService {
	h := HexService{LoadDirections(), NewBattlefield()}
	return h
}

type Coordinate struct {
	Q int
	R int
}

type Cube struct {
	X int
	Y int
	Z int
}

type Tile struct {
	Coordinate Coordinate
	Cost       int
}

type Battlefield map[Coordinate]*Tile

func NewBattlefield() Battlefield {
	b := Battlefield{}

	//Hexagon. Store Hex(q, r) at array[r][q - max(0, N-r)]. Row r size is 2*N+1 - abs(N-r).

	for i := -2; i <= 2; i++ {
		for j := -2; j < 2; j++ {
			b.addDefaultTile(i, j)
		}
	}

	b.addBlockingTile(0, -1)
	b.addBlockingTile(1, 1)
	b.addBlockingTile(2, 1)
	// b.addDefaultTile(4, 0)
	// b.addDefaultTile(5, 0)
	// b.addDefaultTile(6, 0)
	// b.addBlockingTile(2, 1)
	// b.addDefaultTile(3, 1)
	// b.addDefaultTile(4, 1)
	// b.addDefaultTile(5, 1)
	// b.addDefaultTile(6, 1)
	// b.addDefaultTile(1, 2)
	// b.addDefaultTile(2, 2)
	// b.addDefaultTile(3, 2)
	// b.addBlockingTile(4, 2)
	// b.addDefaultTile(5, 2)
	// b.addDefaultTile(6, 2)
	// b.addDefaultTile(0, 3)
	// b.addDefaultTile(1, 3)
	// b.addBlockingTile(2, 3)
	// b.addDefaultTile(3, 3)
	// b.addDefaultTile(4, 3)
	// b.addDefaultTile(5, 3)
	// b.addDefaultTile(6, 3)
	// b.addBlockingTile(0, 4)
	// b.addBlockingTile(1, 4)
	// b.addDefaultTile(2, 4)
	// b.addBlockingTile(3, 4)
	// b.addDefaultTile(4, 4)
	// b.addDefaultTile(5, 4)
	// b.addDefaultTile(0, 5)
	// b.addDefaultTile(1, 5)
	// b.addDefaultTile(2, 5)
	// b.addDefaultTile(3, 5)
	// b.addDefaultTile(4, 5)
	// b.addDefaultTile(0, 6)
	// b.addDefaultTile(1, 6)
	// b.addDefaultTile(2, 6)
	// b.addDefaultTile(3, 6)

	return b
}

func (b Battlefield) addDefaultTile(q int, r int) {
	tile := Tile{Coordinate: Coordinate{Q: q, R: r}, Cost: 1}
	b[Coordinate{Q: q, R: r}] = &tile
}

func (b Battlefield) addBlockingTile(q int, r int) {
	tile := Tile{Coordinate: Coordinate{Q: q, R: r}, Cost: 999}
	b[Coordinate{Q: q, R: r}] = &tile
}

func (b Battlefield) WriteGridCode() {
	var filename string = "gridcode.txt"
	// Create the file if doesn't exist
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	for k, v := range b {
		cube := axialToCube(k)
		msg := fmt.Sprintf("<Hexagon q={%v} r={%v} s={%v} ", cube.X, cube.Z, cube.Y)
		if v.Cost == 999 {
			msg = fmt.Sprint(msg, "fill=\"pat-block\" ")
		}
		msg = fmt.Sprintf("%v><Text>%v, %v, %v</Text></Hexagon>\n", msg, cube.X, cube.Z, cube.Y)
		f.WriteString(msg)
	}
}
