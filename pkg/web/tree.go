
package web

import (
	urlpkg "net/url"
	"regexp"
	"strconv"
	"strings"
)

type patternType int8

const (
	_PATTERN_STATIC    patternType = iota // /home
	_PATTERN_REGEXP
)

// Leaf represents a leaf route information.
type Leaf struct {
	parent *Tree

	typ        patternType
	pattern    string
	rawPattern string // Contains wildcard instead of regexp
	wildcards  []string
	reg        *regexp.Regexp
	optional   bool

	handle Handle
}

// Tree represents a router tree in Macaron.
type Tree struct {
	parent *Tree

	typ        patternType
	pattern    string
	rawPattern string
	wildcards  []string
	reg        *regexp.Regexp

	subtrees []*Tree
	leaves   []*Leaf
}
