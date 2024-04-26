package src

import (
	"Lab_2_algo/general"
	"fmt"
	"sort"
)

type Node struct {
	left, right           *Node
	LeftRange, RightRange int
	sum                   int
}

func AddNode(root *Node, li, ri, val int) *Node {
	if li <= root.LeftRange && ri >= root.RightRange {
		return NodeNew(root.left, root.right, root.LeftRange, root.RightRange, root.sum+val)
	}

	if root.LeftRange >= ri || root.RightRange <= li {
		return root
	}

	toAdd := NodeNew(root.left, root.right, root.LeftRange, root.RightRange, root.sum)

	toAdd.left = AddNode(toAdd.left, li, ri, val)
	toAdd.right = AddNode(toAdd.right, li, ri, val)

	return toAdd
}

func NodeNew(left, right *Node, l, r, sum int) *Node {
	return &Node{left: left,
		right:      right,
		sum:        sum,
		LeftRange:  l,
		RightRange: r,
	}
}

func BuildTree(arr []int, l, r int) *Node {
	if r-l == 1 {
		return NodeNew(nil, nil, l, r, arr[l])
	}
	mid := (r + l) / 2
	left := BuildTree(arr, l, mid)
	right := BuildTree(arr, mid+1, r)
	return NodeNew(left, right, left.LeftRange, right.RightRange, left.sum+right.sum)
}

func BuildPersistentTree(rect *general.Rectangles, CompX, CompY []int) []*Node {
	events := make([]general.Event, len(*rect)*2)
	versions := make([]*Node, 0, len(*rect)*2)
	for _, r := range *rect {
		beginX := general.BinSearch(CompX, r.LB.X)
		beginY := general.BinSearch(CompY, r.LB.Y)
		endY := general.BinSearch(CompY, r.RU.Y+1)
		events = append(events, general.NewEvent(
			beginX,
			beginY,
			endY,
			1,
		))
		events = append(events, general.NewEvent(
			general.BinSearch(CompX, r.RU.X+1),
			general.BinSearch(CompY, r.LB.Y),
			general.BinSearch(CompY, r.RU.Y+1),
			-1,
		))
	}
	sort.Slice(events, func(i, j int) bool { return events[i].X < events[j].X })
	arr := make([]int, len(CompY), len(CompY))
	tree := BuildTree(arr, 0, len(CompY))
	x := events[0].X
	for _, e := range events {
		if x != e.X {
			versions = append(versions, tree)
			x = e.X
		}

		tree = AddNode(tree, e.YStart, e.YEnd, e.IsBegin)
	}
	return versions

}

func SearchInTree(root *Node, val int) int {
	if root != nil {
		mid := (root.LeftRange + root.RightRange) / 2

		if val < mid {
			return root.sum + SearchInTree(root.left, val)
		} else {
			return root.sum + SearchInTree(root.right, val)
		}
	}
	return 0
}

func PersistenceSegmentTree(rects general.Rectangles, ps general.Points) {
	compX, compY := rects.Compress()
	tree := BuildPersistentTree(&rects, compX, compY)
	var res []int
	for _, p := range ps {
		cX := general.BinSearch(compX, p.X)
		cY := general.BinSearch(compY, p.Y)
		if cX == -1 || cY == -1 {
			res = append(res, 0)
			continue
		}
		res = append(res, SearchInTree(tree[cX], cY))
	}
	for _, v := range res {
		fmt.Println(v)
	}

}
