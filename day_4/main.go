package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Hashing("yzbqklnj"))
}

func Hashing(s string) int {
	i := 0
	for {
		value := s + strconv.Itoa(i)
		hash := md5.Sum([]byte(value))
		result := hex.EncodeToString(hash[:])
		fmt.Println(result)
		if result[:5] == "00000" {
			break
		}
		i++

	}

	return i
}
