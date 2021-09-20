package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Conf struct {
	Database struct {
		Connection string
	}
	Converter struct {
		Quality float32
	}
}

func (c *Conf) parse(fileName string) (*Conf, error) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var conf Conf
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", fileName, err)
	}

	return &conf, nil
}

func New(fileName string) (*Conf, error) {
	var c Conf
	return c.parse(fileName)
}
