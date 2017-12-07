package utils

import (
	"io/ioutil"
	"log"
	"regexp"
)

func GetFileNamesByPattern(regexPattern string) ([]string, error) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Panic(err)
	}

	fileNames := make([]string, 0)
	for _, f := range files {
		matched, err := regexp.MatchString(regexPattern, f.Name())

		if err != nil {
			return nil, err
		}

		if matched {
			fileNames = append(fileNames, f.Name())
		}
	}

	return fileNames, nil
}

