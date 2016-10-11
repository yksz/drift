package internal

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

const filename = "config.yaml"

type config struct {
	RootDir string `yaml:"root_dir"`
}

var Conf *config

func init() {
	Conf = &config{}
	if err := load(filename, Conf); err != nil {
		log.Fatal(err)
	}
}

func load(filename string, conf *config) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, conf); err != nil {
		return err
	}
	return nil
}
