package hex

type Direction int32

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

func (h HexService) GetNeighbor(c Coordinate, d Direction) Coordinate {
	dir := h.Directions[d]
	return Coordinate{Q: c.Q + dir.Q, R: c.R + dir.R}
}

type HexService struct {
	Directions Directions
}

func NewHexService() HexService {
	h := HexService{Directions: LoadDirections()}
	return h
}

type Coordinate struct {
	Q int32
	R int32
}

type Tile struct {
	Coordinate Coordinate
	Cost       int32
}

type Battlefield map[Coordinate]*Tile

func NewBattlefield() Battlefield {
	b := Battlefield{}
	b.addDefaultTile(3, 0)
	b.addDefaultTile(4, 0)
	b.addDefaultTile(5, 0)
	b.addDefaultTile(6, 0)
	b.addDefaultTile(2, 1)
	b.addDefaultTile(3, 1)
	b.addDefaultTile(4, 1)
	b.addDefaultTile(5, 1)
	b.addDefaultTile(6, 1)
	b.addDefaultTile(1, 2)
	b.addDefaultTile(2, 2)
	b.addDefaultTile(3, 2)
	b.addDefaultTile(4, 2)
	b.addDefaultTile(5, 2)
	b.addDefaultTile(6, 2)
	b.addDefaultTile(0, 3)
	b.addDefaultTile(1, 3)
	b.addDefaultTile(2, 3)
	b.addDefaultTile(3, 3)
	b.addDefaultTile(4, 3)
	b.addDefaultTile(5, 3)
	b.addDefaultTile(6, 3)
	b.addDefaultTile(0, 4)
	b.addDefaultTile(1, 4)
	b.addDefaultTile(2, 4)
	b.addDefaultTile(3, 4)
	b.addDefaultTile(4, 4)
	b.addDefaultTile(5, 4)
	b.addDefaultTile(0, 5)
	b.addDefaultTile(1, 5)
	b.addDefaultTile(2, 5)
	b.addDefaultTile(3, 5)
	b.addDefaultTile(4, 5)
	b.addDefaultTile(0, 6)
	b.addDefaultTile(1, 6)
	b.addDefaultTile(2, 6)
	b.addDefaultTile(3, 6)

	return b
}

func NewTile(q int32, r int32) Tile {
	return Tile{Coordinate: Coordinate{Q: q, R: r}, Cost: 1}
}

func (b Battlefield) addDefaultTile(q int32, r int32) {
	tile := NewTile(q, r)
	b[Coordinate{Q: q, R: r}] = &tile
}
