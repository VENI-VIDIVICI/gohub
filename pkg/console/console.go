package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

func Success(msg string) {
	colorOut(msg, "green")
}

func Error(msg string) {
	colorOut(msg, "red")
}

func Warining(msg string) {
	colorOut(msg, "yellow")
}

func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

func ExitIf(err error) {
	if err == nil {
		return
	}
	Exit(err.Error())
}
func colorOut(msg, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(msg, color))
}
