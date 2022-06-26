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

const (
	RedactedPassword = "*********"
	DefaultHTTPAddr  = "0.0.0.0"
	Dev              = "development"
	Prod             = "production"
	Test             = "test"
	ApplicationName  = "Grafana"
)

// zoneInfo names environment variable for setting the path to look for the timezone database in go
const zoneInfo = "ZONEINFO"

var (
	// App settings.
	Env = Dev
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
	Raw *ini.File

	// HTTP Server Settings
	CertFile  string
	KeyFile   string
	HTTPAddr  string
	HTTPPort  string
	AppURL    string
	AppSubURL string
	Protocol  Scheme
}
