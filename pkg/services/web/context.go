package web

import (
	"encoding/json"
	"html/template"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// ContextInvoker is an inject.FastInvoker wrapper of func(ctx *Context).
type ContextInvoker func(ctx *Context)

type Context struct {
	Injector
	handler []Handler
	index int

	*Router
	Req *http.Request
	Resp ResponseWriter
	template *template.Template
}
