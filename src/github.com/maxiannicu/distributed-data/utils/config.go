package utils

import "io/ioutil"

func GetConfig(fileName string, object interface{}) error {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}

	return Deserealize(JsonFormat, bytes, object)
}