/*
go-input is package to interact with user input from command line interface.

http://github.com/tcnksm/go-input
*/
package input

import (
	"io"
	"os"
)

var (
	// defaultWriter and defaultReader is default val for UI.Writer
	// and UI.Reader.
	defaultWriter = os.Stdout
	defaultReader = os.Stdin
)

// UI
type UI struct {
	// Writer is where output is written. For example a query
	// to the user will be written here. By default, it's os.Stdout.
	Writer io.Writer

	// Reader is source of input. By default, it's os.Stdin.
	Reader io.Reader
}

// Options
type Options struct {
	// Default is the default value when no thing is input.
	Default string

	// Loop continues to asking user to input until getting
	// valid input.
	Loop bool

	// Required returns error when input is empty.
	Required bool

	// Hide hides user input is prompting console.
	Hide bool

	// Mask hides user input and will be matched by asterisks(*)
	// on the screen.
	Mask bool

	// ValidateFunc is function to validate user input string.
	// By default, it does nothing.
	ValidateFunc ValidateFunc
}

// ValidateFunc is function to validate user input
type ValidateFunc func(string) error

// validateFunc returns ValidateFunc. If it's specifed by
// user it returns it. If not returns default function.
func (o *Options) validateFunc() ValidateFunc {
	if o.ValidateFunc == nil {
		return defaultValidateFunc
	}

	return o.ValidateFunc
}

// defaultValidateFunc does nothing
func defaultValidateFunc(input string) error {
	return nil
}
