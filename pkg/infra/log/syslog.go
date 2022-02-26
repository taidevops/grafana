//go:build !windows && !nacl && !plan9
// +build !windows,!nacl,!plan9

package log

import (
	"log/syslog"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	gokitsyslog "github.com/go-kit/log/syslog"
	"gopkg.in/ini.v1"
)

type SysLogHandler struct {
	syslog *syslog.Writer
	Network string
	Address string
	Facility string
	Tag string
	Format FormatedLogger
	logger log.Logger
}

var selector = func(keyvals ...interface{}) syslog.Priority {

}
