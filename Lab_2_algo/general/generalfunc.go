package general

import (
	"fmt"
)

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func ReadData() (Rectangles, Points, error) {
	var n, m int
	var r Rectangles
	var p Points
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return nil, nil, fmt.Errorf("wrong data format")
	}
	for i := 0; i < n; i++ {
		var x1, x2, y1, y2 int
		if _, err := fmt.Scanf("%d %d %d %d\n", &x1, &y1, &x2, &y2); err != nil {
			return nil, nil, fmt.Errorf("wrong data format")
		}
		LB := Point{X: x1, Y: y1}
		RU := Point{X: x2, Y: y2}
		r = append(r, Rectangle{LB: LB, RU: RU})
	}
	if _, err := fmt.Scanf("%d\n", &m); err != nil {
		return nil, nil, fmt.Errorf("wrong data format")
	}
	for i := 0; i < m; i++ {
		var x, y int
		if _, err := fmt.Scanf("%d %d\n", &x, &y); err != nil {
			return nil, nil, fmt.Errorf("wrong data format")
		}
		p = append(p, Point{X: x, Y: y})
	}
	return r, p, nil
}

func BinSearch(a []int, k int) int {
	l := 0
	r := len(a)
	for l < r {
		m := (r + l) / 2
		if k >= a[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l - 1
}
