package ioJson

import (
	"encoding/json"
	"io/ioutil"
)

func Puts(path string, data interface{}) error {
	json_bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, json_bytes, 0777); err != nil {
		return err
	}

	return nil
}
