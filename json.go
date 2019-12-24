package sac

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	JSONIndent = "  "
)

func (s *Sac) readConfigJSON(path string) error {
	// TODO: Wrap errors betters

	file, err := s.fs.Open(path)
	if err != nil {
		return err
	}

	jsonContent, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonContent, &s.fields)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sac) writeConfigJSON(path string) error {
	content, err := json.MarshalIndent(s.fields, "", JSONIndent)
	if err != nil {
		return err
	}

	file, err := s.fs.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, writePermission)
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	return err
}
