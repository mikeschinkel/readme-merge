package readme_merge

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

const brandKeyword = "merge"

var codeMatcher = regexp.MustCompile("^```")
var headerMatcher = regexp.MustCompile(`^ {0,3}#`)
var linkMatcher = regexp.MustCompile(`^( {0,3}!?)\[([^]]*?)]\(([^)]+?)\)`)

var _ Merger = (*Readme)(nil)

type Readme struct {
	Name     string
	reader   io.Reader
	builder  *strings.Builder
	children map[string]Merger
}

func NewReadme(name string, content string) *Readme {
	reader := strings.NewReader(content)
	return &Readme{
		Name:     name,
		reader:   reader,
		children: make(map[string]Merger),
	}
}

func (rm *Readme) AddChild(m Merger) {
	rm.children[rm.Name] = m
}

func (rm *Readme) GetChild(e string) (m Merger, err error) {
	m, ok := rm.children[e]
	if !ok {
		err = fmt.Errorf("child '%s' not found", e)
	}
	return m, err
}

func (rm *Readme) MergeWithLevel(level int) (string, error) {
	return Merge(rm, rm.Name, level)
}

func (rm *Readme) Reader() (io.Reader, error) {
	return rm.reader, nil
}

func (rm *Readme) CloseReader() {
	rm.reader = nil
}

func (rm *Readme) Builder() *strings.Builder {
	if rm.builder == nil {
		rm.builder = &strings.Builder{}
	}
	return rm.builder
}
