package main

import (
	"fmt"

	"github.com/urbanyeti/go-battle-royale/pkg/hex"
)

func main() {
	h := hex.NewHexService()

	b := hex.NewBattlefield()
	c := hex.Coordinate{Q: 4, R: 1}
	fmt.Println(b[c])

	n := h.GetNeighbor(c, hex.SE)
	fmt.Println(n)
	fmt.Println(b[n])
}
