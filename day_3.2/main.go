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
	MapCountSanta := make(map[[2]int]int)
	MapCountRobo:=make(map[[2]int]int)
	if len(fileLines) == 1 {
		return 2
	}
	var x, y int
	for i := 0; i < len(fileLines); i++ {
		if i==0{
			MapCountSanta[[2]int{0,0}]++
			

			
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
		if i%2==0 {
			MapCountSanta[[2]int{x, y}]++
		}else{
			MapCountRobo[[2]int{x,y}]++
		}
		
	}
	fmt.Println(MapCountSanta,MapCountRobo)
	sum:=(len(MapCountSanta)+len(MapCountRobo))
	return sum
	}

