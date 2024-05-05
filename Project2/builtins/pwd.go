// In builtins/pwd.go

package builtins

import (
	"fmt"
	"io"
	"os"
)

// PrintWorkingDirectory prints the current working directory.
func PrintWorkingDirectory(w io.Writer, args ...string) error {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Print the current working directory to the standard output
	_, err = fmt.Fprintln(w, wd)
	return err
}
