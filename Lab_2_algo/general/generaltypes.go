package general

import "fmt"

type Ordered interface {
	~int | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

type Point struct {
	X, Y int
}

type Rectangle struct {
	LB Point // Left bottom
	RU Point // Right upper
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("%v - %v\n", r.LB, r.RU)

}

type Rectangles []Rectangle
type Points []Point
