package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(CountHouses("input.txt"))
}
func CountHouses(s string) int {
	testFile, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(testFile)
	var fileLines string
	for scanner.Scan() {
		fileLines = scanner.Text()
	}
	sum := 2
	for i := 0; i < len(fileLines); i++ {
		if string(fileLines[i]) == ">" {
			if i > 0 && string(fileLines[i-1]) != "<" {
				sum += 1
			}
		} else if string(fileLines[i]) == "v" {

			if i > 0 && string(fileLines[i-1]) != "^" {
				sum += 1
			}
		} else if string(fileLines[i]) == "^" {
			if i > 0 && string(fileLines[i-1]) != "v" {
				sum += 1
			}
		} else {
			if string(fileLines[i])=="<"{
				if i > 0 && string(fileLines[i-1]) != ">" && string(fileLines[i-2]) != ">" {
					sum += 1
				}
			}
			
		}

	}
	return sum
}
