package readme_merge

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var _ Merger = (*File)(nil)

type File struct {
	Dir     string
	Name    string
	builder *strings.Builder
	file    *os.File
}

func NewInputFile(file string) *File {
	return &File{
		Name: filepath.Base(file),
		Dir:  filepath.Dir(file),
	}
}

func NewOutputFile(dir string) *File {
	var wd string
	var err error
	if !filepath.IsAbs(dir) {
		goto end
	}
	wd, err = os.Getwd()
	if err != nil {
		slog.Error("Getting working directory failed with error: ", err)
		os.Exit(1)
	}
	dir, _ = strings.CutPrefix(dir, wd)
end:
	return &File{Dir: dir, Name: "README.md"}
}

func (f *File) CloseReader() {
	if f.file != nil {
		err := f.file.Close()
		if err != nil {
			slog.Error("Closing file", "filepath", f.Filepath(), "error", err)
		}
	}
	f.file = nil
}

func (f *File) Reader() (_ io.Reader, err error) {
	if f.file == nil {
		f.file, err = os.Open(f.Filepath())
	}
	return f.file, err
}

func (f *File) Builder() *strings.Builder {
	if f.builder == nil {
		f.builder = new(strings.Builder)
	}
	return f.builder
}

func (f *File) GetChild(c string) (m Merger, err error) {
	m = NewInputFile(c)
	_, err = fileExists(f.Filepath())
	return m, err
}

func (f *File) Save(content string) error {
	fp := f.Filepath()
	return os.WriteFile(fp, []byte(content), os.ModePerm)
}

func (f *File) Filepath() string {
	dir, err := os.Getwd()
	if err != nil {
		slog.Error("Unable to get current working directory", "error", err)
		dir = "."
	}
	return filepath.Join(dir, f.Dir, f.Name)
}

func (f *File) Merge() (lines string, err error) {
	return f.MergeWithLevel(0)
}

func (f *File) MergeWithLevel(level int) (lines string, err error) {
	slog.Info("Merging", "file", f.Filepath())
	return Merge(f, f.Dir, level)
}
