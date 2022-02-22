package log

import (
	"fmt"
	"io"
	"os"

	gokitlog "github.com/go-kit/log"
	"github.com/mattn/go-isatty"
)

func FormatedLogger func(w io.Writer) gokitlog.Logger

func getLogFormat(format string) FormatedLogger {
	switch format {
	case "console":
		if isatty.IsTerminal(os.Stdout.Fd()) {
			return func(w io.Writer) gokitlog.Logger {
				return
			}
		}
	}
}
