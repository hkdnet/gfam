package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func CmdList(c *cli.Context) {
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
	fmt.Printf("%v", m)
	return nil, m
}

func getYamlPath() string {
	path := os.Getenv("GOTP_YAML_PATH")
	if path == "" {
		path = "~/.gotp.yml"
	}
	return path
}
