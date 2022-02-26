package build

import (
	"flag"
	"runtime"
)

type BuildOpts struct {
	goarch string
	goos   string
	libc   string

	version   string
}

func BuildOptsFromFlags() BuildOpts {
	opts := BuildOpts{}

	flag.StringVar(&opts.goarch, "goarch", runtime.GOARCH, "GOARCH")
	flag.StringVar(&opts.goos, "goos", runtime.GOOS, "GOOS")
	flag.StringVar(&opts.libc, "libc", "", "LIBC")
	flag.Parse()

	return opts
}
