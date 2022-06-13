package main

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Export struct {
		Width             int `yaml:"width"`
		Height            int `yaml:"height"`
		PreciseIterations int `yaml:"iterations_for_precise_mode"`
	} `yaml:"export"`
}

func (c *Config) initFromFile() {
	yfile, err := ioutil.ReadFile("fractal_explorer_config.yaml")

	if err != nil {
		// create base config
		c.Export.Width = 1920
		c.Export.Height = 1080
		c.Export.PreciseIterations = 200

		file, err := os.Create("fractal_explorer_config.yaml")
		if err != nil {
			panic("Config error.")
		}
		bytes, err := yaml.Marshal(c)
		if err != nil {
			panic("Config error.")
		}
		file.Write(bytes)
		return
	}

	err2 := yaml.Unmarshal(yfile, c)
	if err2 != nil {
		log.Fatal(err2)
	}
}
