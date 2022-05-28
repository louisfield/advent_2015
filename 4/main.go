package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

var puzzleSecretKey = "yzbqklnj"

func main() {

	num := 0

	for checkHexCode(num, 5, "00000") {
		num++
	}
	fmt.Println(num)

	num = 0

	for checkHexCode(num, 6, "000000") {
		num++
	}
	fmt.Println(num)

}

func checkHexCode(num int, maxRange int, constraint string) bool {
	hash := md5.Sum([]byte(puzzleSecretKey + strconv.Itoa(num)))
	firstFive := hex.EncodeToString(hash[:])[0:maxRange]

	if firstFive == constraint {
		return false
	} else {
		return true
	}

}
