package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type RectangularPrism struct {
	length int
	width  int
	height int
}

func main() {

	content, err := ioutil.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\r?\n`)
	contentClean := re.ReplaceAllString(string(content), " ")
	rectangles := createRectangles(strings.Split(contentClean, " "))
	fmt.Println(rectangles)
	result := caluclateAllSurfaceAreas(rectangles)

	fmt.Println(result)

	fmt.Println(calculateAllRibbon(rectangles))
}

func createRectangles(vals []string) []RectangularPrism {
	fmt.Println(vals)
	rectangles := []RectangularPrism{}

	for _, s := range vals {
		intArray := []int{}
		fmt.Println(s)

		for _, i := range strings.Split(s, "x") {
			num, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			intArray = append(intArray, num)
		}
		rectangles = append(rectangles, RectangularPrism{length: intArray[0], width: intArray[1], height: intArray[2]})
	}
	return rectangles
}

func calculateSurfaceArea(l int, w int, h int) int {
	return 2*l*w + 2*w*h + 2*h*l + min([]int{l * w, l * h, w * h})
}

func caluclateAllSurfaceAreas(rectangles []RectangularPrism) int {
	res := 0

	for _, rectangle := range rectangles {
		res = res + calculateSurfaceArea(rectangle.length, rectangle.width, rectangle.height)
	}
	return res
}

func calculateRibbon(s1 int, s2 int, s3 int) int {
	return s1 + s1 + s2 + s2 + (s1 * s2 * s3)
}

func calculateAllRibbon(rectangles []RectangularPrism) int {
	res := 0

	for _, rectangle := range rectangles {
		s := []int{rectangle.length, rectangle.width, rectangle.height}
		sort.Ints(s)
		res = res + calculateRibbon(s[0], s[1], s[2])
	}
	return res
}

func min(vals []int) int {
	var m int
	for i, e := range vals {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}
