package build

import (
	"bytes"
	"flag"
	"fmt"
	// "go/build"
	"log"
	"os"
	// "path/filepath"
	"strconv"
	// "strings"
	"time"
)

const (
	GoOSWindows = "windows"
	GoOSLinux   = "linux"

	ServerBinary = "grafana-server"
)

var binaries = []string{ServerBinary}

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

	packageJSON, err := OpenPackageJSON(wd)
	if err != nil {
		return logError("Error opening package json", err)
	}

	opts.version = packageJSON.Version

	if flag.NArg() == 0 {
		log.Println("Usage: go run build.go build")
		return 1
	}

	for _, cmd := range flag.Args() {
		switch cmd {
		case "setup":
			setup(opts.goos)

		case "build":
			//clean()
			for _, binary := range binaries {
				log.Println("building binaries", cmd)
				// Can't use filepath.Join here because filepath.Join calls filepath.Clean, which removes the `./` from this path, which upsets `go build`
				if err := doBuild(binary, fmt.Sprintf("./pkg/cmd/%s", binary), opts); err != nil {
					log.Println(err)
					return 1
				}
			}
		default:
			log.Println("Unknown command", cmd)
			return 1
		}
	}

	return 0
}

func setup(goos string) {
	args := []string{"install", "-v"}
	if goos == GoOSWindows {
		args = append(args, "-buildmode=exe")
	}
	args = append(args, "./pkg/cmd/grafana-server")
	runPrint("go", args...)
}

func doBuild(binaryName, pkg string, opts BuildOpts) error {
	log.Println("building", binaryName, pkg)
	// libcPart := ""
	// if opts.libc != "" {
	// 	libcPart = fmt.Sprintf("-%s", opts.libc)
	// }
	binary := fmt.Sprintf("./bin/%s", binaryName)

	if opts.goos == GoOSWindows {
		binary += ".exe"
	}

	lf, err := ldflags(opts)
	if err != nil {
		return err
	}

	args := []string{"build", "-ldflags", lf}

	if opts.goos == GoOSWindows {
		// Work around a linking error on Windows: "export ordinal too large"
		args = append(args, "-buildmode=exe")
	}

	args = append(args, "-o", binary)
	args = append(args, pkg)

	runPrint("go", args...)

	return nil
}

func ldflags(opts BuildOpts) (string, error) {
	buildStamp, err := buildStamp()
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	b.WriteString("-w")
	b.WriteString(fmt.Sprintf(" -X main.version=%s", opts.version))
	b.WriteString(fmt.Sprintf(" -X main.commit=%s", getGitSha()))
	b.WriteString(fmt.Sprintf(" -X main.buildstamp=%d", buildStamp))
	b.WriteString(fmt.Sprintf(" -X main.buildBranch=%s", getGitBranch()))
	if v := os.Getenv("LDFLAGS"); v != "" {
		b.WriteString(fmt.Sprintf(" -extldflags \"%s\"", v))
	}

	return b.String(), nil
}

func buildStamp() (int64, error) {
	// use SOURCE_DATE_EPOCH if set.
	if v, ok := os.LookupEnv("SOURCE_DATE_EPOCH"); ok {
		return strconv.ParseInt(v, 10, 64)
	}

	bs, err := runError("git", "show", "-s", "--format=%ct")
	if err != nil {
		return time.Now().Unix(), nil
	}

	return strconv.ParseInt(string(bs), 10, 64)
}
