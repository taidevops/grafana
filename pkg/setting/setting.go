package setting

import (
	"gopkg.in/ini.v1"
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
