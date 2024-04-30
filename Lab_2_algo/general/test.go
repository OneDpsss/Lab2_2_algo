package general

import (
	"math/big"
)

func GenerateTestPoint(n int) Points {
	var points Points
	p1 := 93
	p2 := 94
	for i := 0; i < n; i++ {
		LongX := big.NewInt(int64(p1))
		x := LongX.Exp(LongX, big.NewInt(31), big.NewInt(int64(20*i))).Int64()
		LongY := big.NewInt(int64(p2))
		y := LongX.Exp(LongY, big.NewInt(31), big.NewInt(int64(20*i))).Int64()
		p := Point{X: int(int32(x)), Y: int(int32(y))}
		points = append(points, p)
	}
	return points
}

func GenerateTestRect(n int) Rectangles {
	var rectangles Rectangles
	for i := 0; i < n; i++ {
		LB := Point{X: 10 * i, Y: 10 * i}
		RU := Point{X: 10 * (2*n - i), Y: 10 * (2*n - i)}
		rect := Rectangle{LB: LB, RU: RU}
		rectangles = append(rectangles, rect)
	}
	return rectangles

}

func Benchmark(n int) (Rectangles, Points) {
	testRects := GenerateTestRect(n)
	testPoints := GenerateTestPoint(n)
	return testRects, testPoints

}
