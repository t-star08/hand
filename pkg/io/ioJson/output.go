package ioJson

import (
	"encoding/json"
	"io/ioutil"
)

func Puts(path string, data interface{}) error {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, jsonBytes, 0777); err != nil {
		return err
	}

	return nil
}
