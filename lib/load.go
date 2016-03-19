package lib

import (
	"os"

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Load() (error, map[interface{}]interface{}) {
	path := getYamlPath()
	return loadYaml(path)
}

func loadYaml(path string) (error, map[interface{}]interface{}) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		return err, nil
	}
	return nil, m
}

func getYamlPath() string {
	path := os.Getenv("GFAM_YAML_PATH")
	if path == "" {
		path = os.Getenv("HOME") + "/.gfam.yml"
	}
	return path
}
