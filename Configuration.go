package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Configuration struct {
	values map[string]string
	filename string
}

func newConfiguration(filename string) Configuration {
	configuration := Configuration{filename: filename}
	raw_config, _ := ioutil.ReadFile(filename)
	json.Unmarshal(raw_config, &configuration.values)
	return configuration
}

func (Configuration) getStringValue(key string) string {
	return configuration.values[key]
}
func (Configuration) getIntValue(key string) int {
	value, _ := strconv.Atoi(configuration.values[key])
	return value
}




