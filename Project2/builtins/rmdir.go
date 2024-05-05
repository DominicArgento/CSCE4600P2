// In builtins/rmdir.go

package builtins

import (
	"fmt"
	"os"
)

// RemoveDirectory removes a directory.
func RemoveDirectory(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("rmdir: missing operand")
	}

	// Iterate through each directory name provided as an argument and remove it
	for _, dir := range args {
		err := os.Remove(dir)
		if err != nil {
			return err
		}
	}

	return nil
}
