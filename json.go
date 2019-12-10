package sac

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func (s *Sac) readConfigJSON(path string) error {
	// TODO: Wrap errors betters

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	jsonContent, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonContent, &s.fields)
}
