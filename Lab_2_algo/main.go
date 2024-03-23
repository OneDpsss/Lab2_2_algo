package main

import (
	"Lab_2_algo/general"
	"Lab_2_algo/src"
)

func main() {
	r, p, err := general.ReadData()
	if err != nil {
		return
	}
	src.BruteForce(p, r)

}

//Прямоугольники: {(2,2),(6,8)}, {(5,4),(9,10)}, {(4,0),(11,6)}, {(8,2),(12,12)}
//Точка-ответ:
//(2,2) -> 1
//(12,12) -> 0
//(10,4) -> 2
//(5,5) -> 3
//(2,10) -> 0
//(2,8) -> 0
