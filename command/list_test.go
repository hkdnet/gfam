package command

import (
	"os"
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

func TestLoadYaml(t *testing.T) {
	err, m := loadYaml("../samples/sample.yml")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	want := "hogehoge"
	got := m["hoge"]
	if got != want {
		t.Errorf("want: %v but got: %v", want, got)
	}
}

func TestLoadYamlWithoutFile(t *testing.T) {
	err, _ := loadYaml("../samples/no_such_file_sorry.yml")
	if err == nil {
		t.Errorf("Without file, it should throw an error.")
	}
}

func TestGetYamlPathWithNoEnv(t *testing.T) {
	key := "GFAM_YAML_PATH"
	org := os.Getenv(key)
	os.Setenv(key, "")
	got := getYamlPath()
	want := os.Getenv("HOME") + "/.gfam.yml"
	if got != want {
		t.Errorf("want: %v but got: %v", want, got)
	}
	os.Setenv(key, org)
}

func TestGetYamlPathWithEnv(t *testing.T) {
	key := "GFAM_YAML_PATH"
	org := os.Getenv(key)
	os.Setenv(key, "hoge")
	got := getYamlPath()
	want := "hoge"
	if got != want {
		t.Errorf("want: %v but got: %v", want, got)
	}
	os.Setenv(key, org)
}
