package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(FullWrapperDimension("input.txt"))
}

func FullWrapperDimension(s string) int {
	var fullDimension int
	testFile, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(testFile)
	for scanner.Scan() {
		fileLines := scanner.Text()
		var dimensions []int
		var oneDimension string
		var oneDimensionInt int
		
		for _, ch := range fileLines {
			if ch != 'x' {
				oneDimension += string(ch)
			}
			if ch == 'x' {
				oneDimensionInt, _ = strconv.Atoi(oneDimension)
				dimensions = append(dimensions, oneDimensionInt)
				oneDimension = ""
			}
		}
		if oneDimension != "" {
			oneDimensionInt, _ = strconv.Atoi(oneDimension)
			dimensions = append(dimensions, oneDimensionInt)
		
		}
		
		for l := 0; l < len(dimensions)-2; l++ {
			var surfaceArea []int
			for w := l + 1; w < len(dimensions)-1; w++ {
				for h := w + 1; h < len(dimensions); h++ {
					surfaceArea=append(surfaceArea,dimensions[l]*dimensions[w],dimensions[w]*dimensions[h],dimensions[l]*dimensions[h])
					sort.Ints(surfaceArea)
					fullDimension+= 2*dimensions[l]*dimensions[w] + 2*dimensions[w]*dimensions[h] + 2*dimensions[l]*dimensions[h] + surfaceArea[0]
				}
			}
		}
		

	}
	return fullDimension
}
