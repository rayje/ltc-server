package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Port		string         `json:"port"`
	Name            string         `json:"name"`
}

func getConfig() Config {
	port := flag.String("port", "8080", "The port of the host server")
	name := flag.String("name", "ltclbr-server", "The name of the server")
	configFile := flag.String("config", "config.json", "Location of config file")
	flag.Parse()

	var config Config
	if _, err := os.Stat(*configFile); err == nil {
		file, err := ioutil.ReadFile(*configFile)
		if err != nil {
			fmt.Printf("File error: %v\n", err)
			os.Exit(1)
		}

		err = json.Unmarshal(file, &config)
		if err != nil {
			fmt.Printf("JSON error: %v\n", err)
			os.Exit(1)
		}
	}

	config.setName(*name)
	config.setPort(*port)

	return config
}

func (c *Config) setName(name string) {
	if c.Name == "" {
		c.Name = name
	}
}

func (c *Config) setPort(port string) {
	if c.Port == "" {
		c.Port = port
	}
}
