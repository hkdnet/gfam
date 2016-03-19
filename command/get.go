package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/hkdnet/gfam/lib"
)

func CmdGenerate(c *cli.Context) {
	args := c.Args()
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, "get: please specify of which key you'd like to get a mfa value.")
		os.Exit(2)
  }
	err, yaml := lib.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "LoadError: %v", err)
		os.Exit(1)
	}
	key := args[0]
	secret := getSecretOf(yaml, key)
	if secret == "" {
		fmt.Fprintf(os.Stderr, "LoadError: %v", err)
		os.Exit(1)
	}
	fmt.Println(secret)
}

func getSecretOf(data map[interface{}]interface{}, key string) string {
	ret := data[key]
	if ret != nil {
		return ret.(string)
	} else {
		return ""
	}
}
