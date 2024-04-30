package main

import (
	"Lab_2_algo/general"
	src "Lab_2_algo/solutions"
	"fmt"
	"sort"
	"time"
)

func BruteForceBenchmark(rect general.Rectangles, ps general.Points) time.Duration {
	var a []time.Duration
	for i := 0; i < 20; i++ {
		start := time.Now()
		src.BruteForce(ps, rect)
		a = append(a, time.Since(start))
	}
	sort.Slice(a, func(i, j int) bool {
		return false
	})
	return a[len(a)/2]
}

func CompressBenchmark(rect general.Rectangles, ps general.Points) time.Duration {
	var a []time.Duration
	for i := 0; i < 20; i++ {
		start := time.Now()
		src.CompressCoordinates(rect, ps)
		a = append(a, time.Since(start))
	}
	sort.Slice(a, func(i, j int) bool {
		return false
	})
	return a[len(a)/2]
}

func SegTreeBenchmark(rect general.Rectangles, ps general.Points) time.Duration {
	var a []time.Duration
	for i := 0; i < 20; i++ {
		start := time.Now()
		node := src.BuildPerstitentTree(rect)
		src.SolveTree(node, ps)
		a = append(a, time.Since(start))
	}
	sort.Slice(a, func(i, j int) bool {
		return false
	})
	return a[len(a)/2]
}
func PreprocessingBenchmark(rect general.Rectangles, ps general.Points) (time.Duration, time.Duration) {
	var a, b []time.Duration
	for i := 0; i < 20; i++ {
		start := time.Now()
		rect.CompressedMap()
		a = append(a, time.Since(start))
		start = time.Now()
		src.BuildPerstitentTree(rect)
		b = append(b, time.Since(start))
	}
	return a[len(a)/2], b[len(b)/2]

}

func main() {
	for i := 1; i <= 10000000; i *= 2 {
		TestRect, TestPoints := general.Benchmark(i)
		//BruteTime := BruteForceBenchmark(TestRect, TestPoints)
		//fmt.Print("BruteForce: ", i, "  : ", BruteTime.Nanoseconds(), "  ")
		//CompressTime := CompressBenchmark(TestRect, TestPoints)
		//fmt.Print("Compress: ", i, "  : ", CompressTime.Nanoseconds(), "  ")
		//SegTreeTime := SegTreeBenchmark(TestRect, TestPoints)
		//fmt.Print("SegTree: ", i, "  : ", SegTreeTime.Nanoseconds(), "  ")
		PreCompress, PreSegTree := PreprocessingBenchmark(TestRect, TestPoints)
		fmt.Print("Size:", i, "  Compress:", PreCompress.Nanoseconds(), " SegTree:", PreSegTree.Nanoseconds(), '\n')
		fmt.Println()
	}
}
