package web

import (
	urlpkg "net/url"
	"regexp"
	"strconv"
	"strings"
)

type patternType int8

const (
	_PATTERN_STATIC patternType = iota
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

var wildcardPattern = regexp.MustCompile(`:[a-zA-Z0-9]+`)

func isSpecialRegexp(pattern, regStr string, pos []int) bool {
	return len(pattern) >= pos[1]+len(regStr) && pattern[pos[1]:pos[1]+len(regStr)] == regStr
}

func NewLeaf(parent *Tree, pattern string, handle Handle) *Leaf {
	typ, rawPattern, wildcards, reg := checkPattern(pattern)
	optional := false
	if len(pattern) > 0 && pattern[0] == '?' {
		optional = true
	}
	return &Leaf{parent, typ, pattern, rawPattern, wildcards, reg, optional, handle}
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

