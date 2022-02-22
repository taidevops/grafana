package web

import (
	"net/http"
	"strings"
	"sync"
)

var (
	// Known HTTP methods.
	_HTTP_METHODS = map[string]bool{
		"GET":     true,
		"POST":    true,
		"PUT":     true,
		"DELETE":  true,
		"PATCH":   true,
		"OPTIONS": true,
		"HEAD":    true,
	}
)

// routeMap represents a thread-safe map for route tree.
type routeMap struct {
	lock sync.RWMutex
	routes map[string]map[string]*Leaf
}
