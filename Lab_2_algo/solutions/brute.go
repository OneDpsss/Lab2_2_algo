package src

import (
	"Lab_2_algo/general"
)

func CheckPoint(p general.Point, r *general.Rectangles) int {
	cnt := 0
	for _, k := range *r {
		if k.Belongs(p) {
			cnt++
		}
	}
	return cnt
}

func BruteForce(ps general.Points, r general.Rectangles) {
	for _, p := range ps {
		CheckPoint(p, &r)

	}
}
