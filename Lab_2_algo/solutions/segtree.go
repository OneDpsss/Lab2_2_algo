package src

import (
	"Lab_2_algo/general"
	"sort"
)

var xValues []int
var yValues []int

type Node struct {
	value                 int
	left, right           *Node
	leftIndex, rightIndex int
}

func BinarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := (left + right) / 2
		if array[mid] == target {
			return mid
		}
		if target < array[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func BuildTree(array []int, begin, end int) *Node {
	if begin >= end {
		return &Node{value: array[begin], leftIndex: begin, rightIndex: end}
	}
	mid := (begin + end) / 2
	left := BuildTree(array, begin, mid)
	right := BuildTree(array, mid+1, end)
	return &Node{value: left.value + right.value, left: left, right: right, leftIndex: left.leftIndex, rightIndex: right.rightIndex}
}

func Add(root *Node, begin, end, value int) *Node {
	if begin <= root.leftIndex && root.rightIndex <= end {
		return &Node{value: root.value + value, left: root.left, right: root.right, leftIndex: root.leftIndex, rightIndex: root.rightIndex}
	}
	if root.rightIndex < begin || end < root.leftIndex {
		return root
	}
	node := &Node{value: root.value, left: root.left, right: root.right, leftIndex: root.leftIndex, rightIndex: root.rightIndex}
	node.left = Add(node.left, begin, end, value)
	node.right = Add(node.right, begin, end, value)
	return node
}

func BinSearchTree(root *Node, index int) int {
	if root == nil {
		return 0
	}
	mid := (root.leftIndex + root.rightIndex) / 2
	var value int
	if index <= mid {
		value = BinSearchTree(root.left, index)
	} else {
		value = BinSearchTree(root.right, index)
	}
	return value + root.value
}

func BuildPerstitentTree(rectangles general.Rectangles) []*Node {
	if len(rectangles) == 0 {
		return []*Node{}
	}
	xSet := make(map[int]struct{})
	ySet := make(map[int]struct{})
	for _, rectangle := range rectangles {
		xSet[rectangle.LB.X] = struct{}{}
		xSet[rectangle.RU.X] = struct{}{}
		ySet[rectangle.LB.Y] = struct{}{}
		ySet[rectangle.RU.Y] = struct{}{}
	}

	for x := range xSet {
		xValues = append(xValues, x)
	}
	sort.Ints(xValues)

	for y := range ySet {
		yValues = append(yValues, y)
	}
	sort.Ints(yValues)

	events := make([]general.Event, 2*len(rectangles))
	ptr := 0
	for _, rectangle := range rectangles {
		events[ptr] = general.Event{X: general.BinSearch(xValues, rectangle.LB.X), StartY: general.BinSearch(yValues, rectangle.LB.Y), EndY: general.BinSearch(yValues, rectangle.RU.Y) - 1, State: 1}
		ptr++
		events[ptr] = general.Event{X: general.BinSearch(xValues, rectangle.RU.X), StartY: general.BinSearch(yValues, rectangle.LB.Y), EndY: general.BinSearch(yValues, rectangle.RU.Y) - 1, State: -1}
		ptr++
	}
	sort.Slice(events, func(i, j int) bool { return events[i].X < events[j].X })

	values := make([]int, len(yValues))
	root := BuildTree(values, 0, len(yValues)-1)
	roots := make([]*Node, 2*len(rectangles)+1)
	ptr = 0
	lastX := events[0].X
	for _, event := range events {
		if event.X != lastX {
			roots[ptr] = root
			ptr++
			lastX = event.X
		}
		root = Add(root, event.StartY, event.EndY, event.State)
	}
	return roots
}

func SolveTree(roots []*Node, points general.Points) []int {
	answer := make([]int, len(points))
	ptr := 0
	for _, point := range points {
		xPos := BinarySearch(xValues, point.X)
		yPos := BinarySearch(yValues, point.Y)
		if xPos == -1 || yPos == -1 {
			ptr++
			continue
		}
		answer[ptr] = BinSearchTree(roots[xPos], yPos)
		ptr++
	}
	return answer
}
