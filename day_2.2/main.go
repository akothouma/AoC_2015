package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(FullRibbonDimension("input.txt"))
}

func FullRibbonDimension(s string) int {
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
		var cubicDimension int
		
		for l := 0; l < len(dimensions)-2; l++ {
			var sortInt []int
			for w := l + 1; w < len(dimensions)-1; w++ {
				for h := w + 1; h < len(dimensions); h++ {
					cubicDimension=dimensions[l]*dimensions[w]*dimensions[h]
                    sortInt = append(sortInt, dimensions[l],dimensions[w],dimensions[h])
					sort.Ints(sortInt)
					fullDimension+= 2*sortInt[0] + 2*sortInt[1] + cubicDimension
				}
			}
		}
		

	}
	return fullDimension
}
