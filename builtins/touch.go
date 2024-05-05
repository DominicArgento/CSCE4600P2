// In builtins/touch.go

package builtins

import (
	"os"
)

// TouchFile creates an empty file.
func TouchFile(args ...string) error {
	if len(args) == 0 {
		return nil
	}

	// Iterate through each file name provided as an argument and create an empty file
	for _, file := range args {
		_, err := os.Create(file)
		if err != nil {
			return err
		}
	}

	return nil
}
