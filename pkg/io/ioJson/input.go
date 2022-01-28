package ioJson

import (
	"encoding/json"
	"io/ioutil"
)

func Gets(path string, data interface{}) error {
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	
	if err := json.Unmarshal(jsonFile, &data); err != nil {
		return err
	}

	return nil
}
