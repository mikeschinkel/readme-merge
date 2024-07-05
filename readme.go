package readme_merge

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

var _ Merger = (*Readme)(nil)

type Readme struct {
	Dir      string
	Name     string
	reader   io.Reader
	builder  *strings.Builder
	children map[string]*Readme
	parent   *Readme
}

func NewReadme(name string, content string) *Readme {
	reader := strings.NewReader(content)
	return &Readme{
		Dir:      filepath.Dir(name),
		Name:     filepath.Base(name),
		reader:   reader,
		children: make(map[string]*Readme),
	}
}

func (rm *Readme) Filepath() string {
	return filepath.Join(rm.Dir, rm.Name)
}

func (rm *Readme) AddChild(m Merger) *Readme {
	child := m.(*Readme)
	child.parent = rm
	path := filepath.Join(rm.Dir, child.Name)
	rm.children[path] = child
	return child
}
func (rm *Readme) Root() (root *Readme) {
	root = rm
	for root.parent != nil {
		root = root.parent
	}
	return root
}

func (rm *Readme) GetChild(e string) (m Merger, err error) {
	m, ok := rm.children[e]
	if !ok {
		err = fmt.Errorf("child '%s' not found", e)
	}
	return m, err
}

func (rm *Readme) MergeWithLevel(level int) (string, error) {
	return Merge(rm, rm.Dir, level)
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

func (rm *Readme) Merge() (string, error) {
	return rm.MergeWithLevel(0)
}
