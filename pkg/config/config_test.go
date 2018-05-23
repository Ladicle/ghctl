package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSetConfigDir(t *testing.T) {
	tcs := []struct {
		path  string
		error bool
	}{
		{path: ".", error: false},
		{path: "/invalid/dir", error: true},
	}
	config := Config{}
	for i, tc := range tcs {
		if err := config.SetConfigDir(tc.path); err != nil && !tc.error {
			t.Fatalf("Case #%v: could not set config directory: %v", i, err)
		}
	}
}

func TestSetDefaultConfigDir(t *testing.T) {
	config := Config{}
	if err := config.SetDefaultConfigDir(); err != nil {
		t.Fatalf("could not set default config directory: %v", err)
	}
}

func TestLoadAndSaveConfig(t *testing.T) {
	config := Config{}
	if err := config.LoadConfig(); err != nil {
		t.Fatalf("failed config.LoadConfig() when config file is not exists: %v", err)
	}

	config.ConfigDir = "./fixtures"
	if err := config.LoadConfig(); err != nil {
		t.Fatalf("could not read config: %v", err)
	}
	if got, want := config.Ghctl.CurrentContext, "test2"; got != want {
		t.Fatalf("could not get context(%v) name: got=%v, want=%v", config, got, want)
	}

	dir, err := ioutil.TempDir("", "ghctl")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dir)

	config.ConfigDir = dir
	if err := config.SaveConfig(); err != nil {
		t.Fatalf("could not save config: %v", err)
	}

	path := getConfigFilePath(dir)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("could not write config file to %v", path)
	}
}

func TestContext(t *testing.T) {
	config := Config{}
	c1 := Context{Name: "test1", AccessToken: "dummy_token1"}
	c2 := Context{Name: "test2", AccessToken: "dummy_token2"}

	for i, c := range []Context{c1, c2} {
		if err := config.RegisterContext(c); err != nil {
			t.Fatalf("could not register context #%v: %v", i+1, err)
		}
	}

	if err := config.SetCurrentContext(c2.Name); err != nil {
		t.Fatalf("could not set current context: %v", err)
	}
	if got, want := config.GetCurrentContext(), c2; *got != want {
		t.Fatalf("could not valid current context: got=%v, want=%v", *got, want)
	}
	if got, want := config.GetContext(c1.Name), c1; *got != want {
		t.Fatalf("could not get expected context: got=%v, want=%v", *got, want)
	}
}
