package src

import (
	"Lab_2_algo/general"
	"fmt"
)

func CompressCoordinates(r general.Rectangles, p general.Points) {
	CompX, CompY, mp := r.CompressedMap()
	for i := range mp {
		for j := range mp[i] {
			fmt.Printf("%d ", mp[i][j])
		}
		fmt.Println()
	}
	for i := range p {
		X := general.BinSearch(CompX, p[i].X)
		Y := general.BinSearch(CompY, p[i].Y)
		if X == -1 || Y == -1 {
			fmt.Println(p[i], 0)
		} else {
			fmt.Println(p[i], mp[Y][X])
		}
	}
	return
}
