package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestMkDirAllIfNotExist(t *testing.T) {
	path, err := ioutil.TempDir("", "ghctl-test")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(path)

	unextPath := filepath.Join(path, "a", "b", "c", "file")
	if err := MkDirAllIfNotExist(unextPath); err != nil {
		t.Fatalf("could not create directories: %v", err)
	}
	dir := filepath.Dir(unextPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Fatalf("%v is not exists", dir)
	}
}
