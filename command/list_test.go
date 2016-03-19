package command

import (
	"testing"
)

func TestCmdList(t *testing.T) {
	// Write your code here
}

func TestYamlToList(t *testing.T) {
	m := make(map[interface{}]interface{})
	m["hoge"] = "hogehoge"
	m["fuga"] = "fugafuga"
	got := yamlToList(m)
	want := "hoge\nfuga"
	if got != want {
		t.Errorf("want: %v but got: %v", want, got)
	}
}
