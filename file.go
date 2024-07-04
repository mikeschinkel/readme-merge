package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Level   int
	Dir     string
	Name    string
	Builder strings.Builder
}

func (f File) Save(content string) error {
	fp := f.Filepath()
	return os.WriteFile(fp, []byte(content), os.ModePerm)
}

func NewInputFile(file string, level int) *File {
	return &File{
		Name:  filepath.Base(file),
		Dir:   filepath.Dir(file),
		Level: level,
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

func (f File) Filepath() string {
	dir, err := os.Getwd()
	if err != nil {
		slog.Error("Unable to get current working directory", "error", err)
		dir = "."
	}
	return filepath.Join(dir, f.Dir, f.Name)
}
