package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	textfile, _ := os.Open("input.txt")
	scanner1 := bufio.NewScanner(textfile)

	var fileLines string
	for scanner1.Scan() {
		fileLines = scanner1.Text()
	}
	fmt.Println(ReGifted(fileLines))
}

func ReGifted(s string) int {
	MapSanta := make(map[[2]int]int)
	MapRobo := make(map[[2]int]int)
	var x, y, x1, y1 int
	if len(s) == 1 {
		return 2
	}

	for i := 0; i < len(s)-1; i += 2 {
		if i == 0 {
			MapSanta[[2]int{0, 0}]++
		}
		fmt.Println(string(s[i]))
		switch string(s[i]) {

		case "^":
			y++
		case ">":
			x++
		case "v":
			y--
		case "<":
			x--
		}
		MapSanta[[2]int{x, y}]++
	}

	for j := 1; j < len(s); j += 2 {
		if j == 1 {
			MapSanta[[2]int{0, 0}]++
		}
		fmt.Println(string(s[j]))
		switch string(s[j]) {
		case "^":
			y1++
		case ">":
			x1++
		case "v":
			y1--
		case "<":
			x1--
		}
		MapSanta[[2]int{x1, y1}]++
	}
	fmt.Println(MapSanta, MapRobo)
	return len(MapSanta)
}

// func Unique(m1, m2 map[[2]int]int) (map[[2]int]int, map[[2]int]int) {
// 	uniqueMap := make(map[[2]int]int)
// 	commonMap := make(map[[2]int]int)
// 	for key, value := range m1 {
// 		if _, ok := m2[key]; ok {
// 			commonMap[key] = value
// 		} else {
// 			uniqueMap[key] = value
// 		}
// 	}
// 	for key, value := range m2 {
// 		if _, ok := m1[key]; ok {
// 			commonMap[key] = value
// 		} else {
// 			uniqueMap[key] = value
// 		}
// 	}
// 	fmt.Println(uniqueMap,"unique\n",commonMap)
// 	return uniqueMap, commonMap
// }

// var x, y int

// func Count(s string) map[[2]int]int {
// 	MapCount := make(map[[2]int]int)
// 	MapCount[[2]int{0, 0}]++
// 	switch string(s) {
// 	case "^":
// 		y++
// 	case ">":
// 		x++
// 	case "v":
// 		y--
// 	case "<":
// 		x--
// 	}
// 	MapCount[[2]int{x, y}]++
// 	return MapCount
// }
