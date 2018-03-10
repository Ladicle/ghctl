package util

import (
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// LoadYAML reads data from YAML file and expand it to `out`.
func LoadYAML(path string, out interface{}) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	y, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(y, out)
}

// WriteYAML writes `in` to YAML file.
func WriteYAML(path string, in interface{}) error {
	y, err := yaml.Marshal(in)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, y, 0644)
}

// MkDirAllIfNotExist makes all directories if it not exist.
func MkDirAllIfNotExist(path string) (err error) {
	dir := filepath.Dir(path)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}
	return err
}
