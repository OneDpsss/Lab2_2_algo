## Лабораторная работа №2: Проверка принадлежности точек прямоугольникам

### Описание проекта

**Язык программирования:** Go (golang) 1.22.2

**Фреймворк:** go-echarts

### Структура проекта

- **genral:** Пакет с необходимыми для проекта функциями и типами
  - **generalfuncs.go:** Общие для всех алгоритмов функции
  - **generaltypes.go:** Общие структуры и методы для выполнения работы
- **solutions:** Пакет для запуска реализаций
  - brute.go: Реализация решения в лоб
  - compress.go: Реализация через карту
  - segtree.go: Реализация через персистентное дерево отрезков
  - test.go: Генерация тестов для бенчмарков

### Входные данные

Программа принимает входные данные в виде N прямоугольников и M точек, заданных координатами x и y.

```go
type Point struct {
	X, Y int
}

type Rectangle struct {
	LB Point // Left bottom
	RU Point // Right upper
}
```

### Решение №1

- **Алгоритмическая сложность по времени:** O(M * N)
- **Алгоритмическая сложность по памяти:** O(1)

Этот алгоритм заключается в полном переборе всех M точек и проверке принадлежности каждой из них каждому из N прямоугольников.

```go
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
```

### Решение №2

- **Алгоритмическая сложность по времени:** O(N^3) на построение карты и O(M * logN) на поиск ответа
- **Алгоритмическая сложность по памяти:** O(set(x) * set(y))

Этот алгоритм сначала предварительно сжимает координаты прямоугольников по обеим осям, затем строит карту и использует бинарный поиск для поиска ответа.
```go
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



```

### Решение №3

- **Алгоритмическая сложность по времени:** O(N*logN) на построение карты и O(M * logN) на поиск ответа
- **Алгоритмическая сложность по памяти:** O(N^2)

Этот алгоритм строит персистентное дерево отрезков для оптимизации предобработки из решения №2 и ускорения поиска ответа.

```go
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


```
### Тесты

Тестирование проводилось на данных, где число прямоугольников и точек равно 2^n, 1 <= n <= x.

```go
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


```
### Графики

#### Время выполнения

-<img width="434" alt="image" src="https://github.com/OneDpsss/Lab2_2_algo/assets/108849165/e2a54e45-b18d-420a-b304-54f1f76a5c2e">


#### Время на предобработку

- 
- 

#### Общее время

- 
- 

### Общий вывод

Использование алгоритма №2 (через карту) имеет смысл на малых данных, но начиная с определенного размера данных становится неэффективным из-за большой сложности. Алгоритм №3 (через персистентное дерево отрезков) наиболее эффективен на больших объемах данных, хотя требует больше памяти для предобработки. Алгоритм №1 (в лоб) подходит для небольших нагрузок, но на больших объемах данных становится неэффективным из-за своей сложности.
