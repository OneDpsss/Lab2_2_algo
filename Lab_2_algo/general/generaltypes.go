package general

import (
	"sort"
)

type Point struct {
	X, Y int
}

type Rectangle struct {
	LB Point // Left bottom
	RU Point // Right upper
}

func (r Rectangle) Belongs(p Point) bool {
	return (p.X >= r.LB.X && p.X < r.RU.X) && (p.Y >= r.LB.Y && p.Y < r.RU.Y)
}

func (r *Rectangles) CompressedMap() ([]int, []int, [][]int) {
	CompX, CompY := r.Compress()
	mp := make([][]int, len(CompY))
	for i := range mp {
		mp[i] = make([]int, len(CompX))
	}
	for _, rect := range *r {
		lx := BinSearch(CompX, rect.LB.X)
		ly := BinSearch(CompY, rect.LB.Y)
		rx := BinSearch(CompX, rect.RU.X+1)
		ry := BinSearch(CompY, rect.RU.Y+1)
		//fmt.Println(lx, ly, rx, ry)
		for i := ly; i < ry; i++ {
			for j := lx; j < rx; j++ {
				mp[i][j]++
			}
		}
	}
	return CompY, CompX, mp
}

func (r Rectangles) Compress() ([]int, []int) {
	CompXset := make(map[int]struct{}, len(r)*2)
	CompYset := make(map[int]struct{}, len(r)*2)
	for _, rect := range r {
		CompXset[rect.LB.X] = struct{}{}
		CompXset[rect.RU.X+1] = struct{}{}
		CompYset[rect.LB.Y] = struct{}{}
		CompYset[rect.RU.Y+1] = struct{}{}
	}
	var CompX = make([]int, 0, len(CompXset))
	var CompY = make([]int, 0, len(CompYset))
	for k := range CompXset {
		CompX = append(CompX, k)
	}
	for k := range CompYset {
		CompY = append(CompY, k)
	}
	sort.Ints(CompX)
	sort.Ints(CompY)
	return CompX, CompY
}

type Rectangles []Rectangle
type Points []Point

type Event struct {
	X, StartY, EndY, State int
}

func NewEvent(x, yStart, yEnd, isBegin int) Event {
	return Event{
		X:      x,
		StartY: yStart,
		EndY:   yEnd,
		State:  isBegin,
	}
}
