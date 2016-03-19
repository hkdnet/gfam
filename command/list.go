package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/hkdnet/gfam/lib"
)

func CmdList(c *cli.Context) {
	err, yaml := lib.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "LoadError: %v", err)
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
