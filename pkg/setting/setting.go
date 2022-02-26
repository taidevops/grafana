package setting

import (
	"gopkg.in/ini.v1"
)

type Scheme string

const (
	HTTPScheme   Scheme = "http"
	HTTPSScheme  Scheme = "https"
	HTTP2Scheme  Scheme = "h2"
	SocketScheme Scheme = "socket"
)

var (
	// build
	BuildVersion string
	BuildCommit  string
	BuildBranch  string
	BuildStamp   int64

)

// TODO move all global vars to this struct
type Cfg struct {
	Raw    *ini.File
	Logger log.Logger

	// HTTP Server Settings
	CertFile         string
	KeyFile          string
	HTTPAddr         string
	HTTPPort         string
	AppURL           string
	AppSubURL        string
}
