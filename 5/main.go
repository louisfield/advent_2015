package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var badStrings = []string{"ab", "cd", "pq", "xy"}
var vowels = "aeiou"
var vowelsRegex = `[aeiou]`

func main() {
	content, err := ioutil.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\r?\n`)
	contentClean := strings.Split(re.ReplaceAllString(string(content), " "), " ")
	task_part_1(contentClean)
	task_part_2(contentClean)
}

func task_part_1(contentClean []string) {
	niceAmount := 0
	badAmount := 0

	for _, v := range contentClean {

		stringBad := true
		splitStrings := strings.Split(v, "")
		for i, l := range splitStrings[:len(splitStrings)-1] {
			if l == splitStrings[i+1] {
				stringBad = false
				break
			}
		}

		re := regexp.MustCompile(vowelsRegex)
		if len(re.FindAllString(v, -1)) < 3 {
			stringBad = true
		}

		for _, bV := range badStrings {
			if strings.Contains(v, bV) {
				stringBad = true
			}
		}

		if stringBad {
			badAmount++
		} else {
			niceAmount++
		}

	}

	fmt.Println(niceAmount)
	fmt.Println(badAmount)
}

func task_part_2(contentClean []string) {
	niceAmount := 0
	badAmount := 0

	for _, v := range contentClean {

		stringBad := false
		stringGood := false
		stringGood2 := false
		stringGood3 := false
		splitStrings := strings.Split(v, "")
		for i, l := range splitStrings[:len(splitStrings)-1] {
			if i < len(splitStrings)-2 {
				if l == splitStrings[i+1] && l == splitStrings[i+2] {
					stringBad = true
				}
				if l == splitStrings[i+2] && l != splitStrings[i+1] {
					stringGood2 = true
				}

			}
			if strings.Contains(strings.Replace(v, l, "", 1), l) {
				stringGood3 = true
			}
			if strings.Contains(strings.Replace(v, l+splitStrings[i+1], "", 1), l+splitStrings[i+1]) {
				stringGood = true
			}
		}

		if !stringBad && stringGood && stringGood2 && stringGood3 {
			niceAmount++
		} else {
			badAmount++
		}

	}

	fmt.Println(niceAmount)
	fmt.Println(badAmount)
}
