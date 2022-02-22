package term

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	gokitlog "github.com/go-kit/log"
	"github.com/grafana/grafana/pkg/infra/log/level"
)

const (
	timeFormat     = "2006-01-02T15:04:05-0700"
	termTimeFormat = "01-02|15:04:05"
	floatFormat    = 'f'
	termMsgJust    = 40
	errorKey       = "LOG15_ERROR"
)

type terminalLogger struct {
	w io.Writer
}

