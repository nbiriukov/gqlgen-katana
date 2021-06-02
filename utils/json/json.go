package json

import (
	jsonEncode "encoding/json"
	"io/ioutil"
)

func ImportFromFile(path string) (interface{}, error) {
	byteArray, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = jsonEncode.Unmarshal(byteArray, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}