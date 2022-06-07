package internal

import (
	"io/ioutil"
)

func ExtractFile(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)
	return string(content), err

}
