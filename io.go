package vivado

import (
	"io"
	"os"
)

var (
	stdout io.Writer = os.Stdout
)

func SetOutput(w io.Writer) {
	stdout = w
}
