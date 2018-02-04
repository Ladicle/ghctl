package util

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// HandleCmdError handles command processing error.
// If err is not nil, output it to stderr.
func HandleCmdError(e error, errOut io.Writer) {
	if e != nil {
		fmt.Fprintf(errOut, "%v\n", e)
		os.Exit(-1)
	}
}

// GetPrettyOutput returns pretty marshaled data.
// This function supports yaml and json format.
func GetPrettyOutput(format string, v interface{}) ([]byte, error) {
	if format == "yaml" {
		return yaml.Marshal(v)
	}
	if format == "json" {
		return json.MarshalIndent(v, "", "  ")
	}
	return []byte{}, fmt.Errorf("%s is unknown format.", format)
}
