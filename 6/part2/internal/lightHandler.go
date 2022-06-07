package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func performLightTask(numArray []int, value bool, lights *[1000][1000]int, toggle bool) {

	for i := numArray[0]; i <= numArray[2]; i++ {
		for j := numArray[1]; j <= numArray[3]; j++ {
			if toggle {
				(*lights)[i][j] = (*lights)[i][j] + 2
			} else if value {
				(*lights)[i][j] = (*lights)[i][j] + 1
			} else {
				if (*lights)[i][j] > 0 {
					(*lights)[i][j] = (*lights)[i][j] - 1
				}
			}

		}
	}
}

func performContentCycle(content []string, lights *[1000][1000]int) {
	re := regexp.MustCompile(`\d+`)

	for _, v := range content {

		vals := re.FindAllString(v, -1)
		numArray := []int{}
		for _, j := range vals {
			j, err := strconv.Atoi(j)
			if err != nil {
				panic(err)
			}
			numArray = append(numArray, j)
		}
		fmt.Println("Performing: " + v)
		if strings.Contains(v, "off") {
			performLightTask(numArray, false, lights, false)
			// for loop through first two and then through other two in for loop to turn on/off
		} else if strings.Contains(v, "on") {
			performLightTask(numArray, true, lights, false)
		} else {
			performLightTask(numArray, true, lights, true)
		}

	}
}
func FirstTask(content []string, lights [1000][1000]int) [1000][1000]int {

	performContentCycle(content, &lights)

	lightsOn := 0
	for i := 0; i < len(lights); i++ {
		for j := 0; j < len(lights[i]); j++ {
			lightsOn += lights[i][j]
		}
	}

	fmt.Println(lightsOn)
	return lights
}
