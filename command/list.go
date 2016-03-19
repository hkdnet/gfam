package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func CmdList(c *cli.Context) {
	path := getYamlPath()
	err, yaml := loadYaml(path)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	fmt.Println(yamlToList(yaml))
}

func yamlToList(yaml map[interface{}]interface{}) string {
	keys := make([]string, len(yaml))
	i := 0
	for k := range yaml {
		keys[i] = k.(string)
		i++
	}
	return strings.Join(keys, "\n")
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
