package version

import (
	"fmt"
	"io"
	"os"
)

func FprintVersion(w io.Writer) {
	fmt.Fprintln(w, os.Args[0], Package, Version)
}

func PrintVersion() {
	FprintVersion(os.Stdout)
}
