package sac

import (
	"io/ioutil"
	"os"
        "fmt"

	"gopkg.in/yaml.v3"
)

func (s *Sac) readConfigYAML(path string) error {
	// TODO: Wrap errors betters

	file, err := s.fs.Open(path)
	if err != nil {
		return err
	}

	yamlContent, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlContent, s.fields)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sac) writeConfigYAML(path string) error {
	content, err := yaml.Marshal(s.fields)
	if err != nil {
		return err
	}

        fmt.Println(string(content))

	file, err := s.fs.OpenFile(path, os.O_CREATE|os.O_WRONLY, writePermission)
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	return err
}
