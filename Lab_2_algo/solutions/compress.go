package src

import (
	"Lab_2_algo/general"
)

func CompressCoordinates(r general.Rectangles, p general.Points) {
	CompX, CompY, _ := r.CompressedMap()

	for i := range p {
		X := general.BinSearch(CompX, p[i].X)
		Y := general.BinSearch(CompY, p[i].Y)
		if X == -1 || Y == -1 {
			//fmt.Println(p[i], 0)
			continue
		} else {
			//fmt.Println(p[i], mp[Y][X])
			continue
		}
	}
	return
}
