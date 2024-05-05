// In builtins/echo.go

package builtins

import (
	"fmt"
	"io"
	"strings"
)

// Echo prints its arguments to the standard output.
func Echo(w io.Writer, args ...string) error {
	// Join all arguments into a single string with spaces between them
	message := strings.Join(args, " ")

	// Print the message to the standard output
	_, err := fmt.Fprintln(w, message)
	return err
}
