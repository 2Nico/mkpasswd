package src

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
)

func HandleInputError(Len uint64, numLen uint64, charLen uint64) {
	if numLen > Len || charLen > Len {
		fmt.Println("[" + color.Ize(color.Red, "ERROR") + "] The default length of password is 16, and -n and -c must be smaller than it.")
		os.Exit(1)
	}
}
