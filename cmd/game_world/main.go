package main

import (
	"github.com/urbanyeti/go-battle-royale/pkg/hex"
)

func main() {
	h := hex.NewHexService()
	hex.FindPath(h, hex.Coordinate{Q: 2, R: 5}, hex.Coordinate{Q: 3, R: 2})
	// c := hex.Coordinate{Q: 4, R: 1}
	// t := h.Battlefield[c]
	// fmt.Println(t)

	// n := h.GetNeighbor(t, hex.SE)
	// fmt.Println(n)

}
