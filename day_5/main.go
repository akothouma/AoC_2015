package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	textfile, _ := os.Open("input.txt")
	scanner1 := bufio.NewScanner(textfile)
	sum := 0
	var fileLines string
	for scanner1.Scan() {
		fileLines = scanner1.Text()
		if ValidString(fileLines) {
			sum++
		}
	}
	fmt.Println(sum)
}

func ValidString(s string) bool {

	Excluded := []string{"ab", "cd", "pq", "xy"}

	for j := 0; j < len(Excluded); j++ {
		if strings.Contains(s, Excluded[j]) {
			return false
		}
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "a" || string(s[i]) == "e" || string(s[i]) == "i" || string(s[i]) == "o" || string(s[i]) == "u" {
			count++
		}
	}

	if count >= 3 {
		for i := 0; i < len(s); i++ {
			if i < len(s)-1 && s[i] == s[i+1] {
				return true
			}
		}

	}

	return false
}
