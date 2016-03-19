package command

import (
	"os"
	"testing"
)

func TestCmdList(t *testing.T) {
	// Write your code here
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

func TestGetYamlPathWithNoEnv(t *testing.T) {
	key := "GOTP_YAML_PATH"
	org := os.Getenv(key)
	os.Setenv(key, "")
	got := getYamlPath()
	want := "~/.gotp.yml"
	if got != want {
		t.Errorf("want: %v but got: %v", want, got)
	}
	os.Setenv(key, org)
}

func TestGetYamlPathWithEnv(t *testing.T) {
	key := "GOTP_YAML_PATH"
	org := os.Getenv(key)
	os.Setenv(key, "hoge")
	got := getYamlPath()
	want := "hoge"
	if got != want {
		t.Errorf("want: %v but got: %v", want, got)
	}
	os.Setenv(key, org)
}
