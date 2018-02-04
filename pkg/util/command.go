package util

import (
	"fmt"
	"io"
	"os"
)

// HandleCmdError handles command processing error.
// If err is not nil, output it to stderr.
func HandleCmdError(e error, errOut io.Writer) {
	if e != nil {
		fmt.Fprintf(errOut, "%v\n", e)
		os.Exit(-1)
	}
}
