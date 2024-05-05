// In builtins/mkdir.go

package builtins

import (
	"fmt"
	"os"
)

// MakeDirectory creates a new directory.
func MakeDirectory(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("mkdir: missing operand")
	}

	// Iterate through each directory name provided as an argument and create it
	for _, dir := range args {
		err := os.Mkdir(dir, 0755) // 0755 is the default permission mode for directories
		if err != nil {
			return err
		}
	}

	return nil
}
