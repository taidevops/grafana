package build

import (
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	GoOSLinux   = "linux"
)

func logError(message string, err error) int {
	log.Println(message, err)

	return 1
}

// RunCmd runs the build command and returns the exit code
func RunCmd() int {
	opts := BuildOptsFromFlags()

	wd, err := os.Getwd()
	if err != nil {
		return logError("Error getting working directory", err)
	}


}