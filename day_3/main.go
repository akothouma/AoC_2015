package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(ReGifted("input.txt"))
}

func ReGifted(s string) int {
	textfile, _ := os.Open(s)
	scanner1 := bufio.NewScanner(textfile)

	var fileLines string
	for scanner1.Scan() {
		fileLines = scanner1.Text()
	}
	MapCount := make(map[[2]int]int)
	if len(fileLines) == 1 {
		return 2
	}
	var x, y int
	for i := 0; i < len(fileLines); i++ {
		if i == 0 {
			MapCount[[2]int{x, y}]++
		}
		switch string(fileLines[i]) {
		case "^":
			y++
		case ">":
			x++
		case "v":
			y--
		case "<":
			x--
		}
		MapCount[[2]int{x, y}]++
	}
	return len(MapCount)
	// fmt.Println(string(fileLines[0]), string(fileLines[1]))
	// sum := 1

	// if len(fileLines) == 1 {
	// 	sum = sum + 1
	// 	return sum
	// }
	// for i := 0; i < len(fileLines)-1; i++ {
	// 	if string(fileLines[i]) == "^" {
	// 		if string(fileLines[i+1]) == ">" || string(fileLines[i+1]) == "<" || string(fileLines[i+1]) == "^"|| string(fileLines[i+1]) == "v" {
	// 			sum += 1
	// 		}

	// 	}else if string(fileLines[i]) == ">" {
	// 		if string(fileLines[i+1]) == "v" || string(fileLines[i+1]) == ">" || string(fileLines[i+1]) == "^" || string(fileLines[i+1]) == "<" {
	// 			sum += 1
	// 		}

	// 	} else if string(fileLines[i]) == "v" {
	// 		if string(fileLines[i+1]) == "v" || string(fileLines[i+1]) == "<" || string(fileLines[i+1]) == ">" {
	// 			sum += 1
	// 		}
	// 		if string(fileLines[i+1])=="^"{
	// 			sum=sum-1
	// 		}
	// 	} else {
	// 		if string(fileLines[i+1]) == "v" || string(fileLines[i+1]) == "<" || string(fileLines[i+1]) == "^" {
	// 			sum += 1
	// 		}
	// 	}
	// }

	// return sum
}
