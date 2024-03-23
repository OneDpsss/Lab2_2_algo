package general

import (
	"fmt"
)

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

func (r Rectangle) Belongs(p Point) bool {
	return (p.X >= r.LB.X && p.X < r.RU.X) && (p.Y >= r.LB.Y && p.Y < r.RU.Y)
}
