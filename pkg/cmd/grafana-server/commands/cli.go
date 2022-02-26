package commands

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/setting"
)

type ServerOptions struct {
	Version     string
	Commit      string
	BuildBranch string
	BuildStamp  string
}

type exitWithCode struct {
	reason string
	code   int
}

var serverFs = flag.NewFlagSet("server", flag.ContinueOnError)

func (e exitWithCode) Error() string {
	return e.reason
}

func RunServer(opt ServerOptions) int {
	var (
		//homePath   = serverFs.String("homepath", "", "path to grafana install/home path, defaults to working directory")

		v           = serverFs.Bool("v", false, "prints current version and exits")
		vv          = serverFs.Bool("vv", false, "prints current version, all dependencies and exits")
	)

	if err := serverFs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	if *v || *vv {
		fmt.Printf("Version %s (commit: %s, branch: %s)\n", opt.Version, opt.Commit, opt.BuildBranch)
		if *vv {
			fmt.Println("Dependencies:")

		}
		return 0
	}

	ctx := context.Background()

	go listenToSystemSignals(ctx)

	return 0
}

func executeServer(opt ServerOptions) error {

	buildstampInt64, err := strconv.ParseInt(opt.BuildStamp, 10, 64)
	if err != nil || buildstampInt64 == 0 {
		buildstampInt64 = time.Now().Unix()
	}

	setting.BuildVersion = opt.Version
	setting.BuildCommit = opt.Commit
	setting.BuildStamp = buildstampInt64
	setting.BuildBranch = opt.BuildBranch


}


func listenToSystemSignals(ctx context.Context) {
	signalChan := make(chan os.Signal, 1)
	sighupChan := make(chan os.Signal, 1)

	signal.Notify(sighupChan, syscall.SIGHUP)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sighupChan:
			if err := log.Reload(); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to reload loggers: %s\n", err)
			}
		case <-signalChan:
			_, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			return
		}
	}
}
