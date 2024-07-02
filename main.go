package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func mustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		slog.Warn("Error occurred when calling mustClose()", "error", err)
	}
}

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
	return &File{
		Name: "README.md",
		Dir:  dir,
	}
}

func (f File) Filepath() string {
	dir, err := os.Getwd()
	if err != nil {
		slog.Error("Unable to get current working directory", "error", err)
		dir = "."
	}
	return filepath.Join(dir, f.Dir, f.Name)
}

var headerMatcher = regexp.MustCompile(`^#`)
var includeMatcher = regexp.MustCompile(`^\[include]\(([^)]+?)\)$`)

func (f File) Parse() (lines string, err error) {
	var scanner *bufio.Scanner
	var line string
	var file *os.File

	slog.Info("Reading", "file", f.Filepath())
	file, err = os.Open(f.Filepath())
	if err != nil {
		err = fmt.Errorf("opening index file; %w", err)
		goto end
	}
	defer mustClose(file)
	scanner = bufio.NewScanner(file)

	// Debug: Check if the scanner is correctly initialized
	if scanner == nil {
		err = fmt.Errorf("scanner not initialized")
		goto end
	}

	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			f.Builder.WriteByte('\n')
			continue
		}
		matches := includeMatcher.FindStringSubmatch(line)
		if len(matches) != 0 {
			file := filepath.Join(f.Dir, matches[1])
			nf := NewInputFile(file, f.Level+1)
			lines, err = nf.Parse()
			f.Builder.WriteString(lines)
			continue
		}
		if f.Level != 0 && headerMatcher.MatchString(line) {
			line = strings.Repeat("#", f.Level) + line
		}
		f.Builder.WriteString(line)
		f.Builder.WriteByte('\n')
	}

	if err = scanner.Err(); err != nil {
		err = fmt.Errorf("scanning index file; %w", err)
		goto end
	}

	lines = f.Builder.String()
end:
	return lines, err
}

func main() {
	indexFile, readmeDir := parseArgs()
	slog.Info("Initializing")
	file := NewInputFile(indexFile, 0)
	slog.Info("Parsing", "index_file", indexFile)
	readme, err := file.Parse()
	if err != nil {
		slog.Error("Parsing index file", "index_file", indexFile, "error", err)
		os.Exit(2)
	}
	file = NewOutputFile(readmeDir)
	slog.Info("Saving", "readme_file", readmeDir)
	err = file.Save(readme)
	if err != nil {
		slog.Error("Saving README.md", "readme_file", readmeDir, "error", err)
		os.Exit(3)
	}
	slog.Info("Done")
}

func parseArgs() (indexFile, readmePath string) {
	switch {
	case len(os.Args) < 3:
		goto err
	case len(os.Args[1]) == 0:
		goto err
	case len(os.Args[2]) == 0:
		goto err
	case !fileMustExist(os.Args[1]):
		goto err
	case !dirMustExist(os.Args[2]):
		goto err
	}
	indexFile = os.Args[1]
	readmePath = os.Args[2]
	goto end
err:
	_, _ = fmt.Fprintln(os.Stderr, "Usage: readme-merge <index_file> <readme_path>")
	os.Exit(1)
end:
	return indexFile, readmePath
}

func fileMustExist(file string) (exists bool) {
	info, err := os.Stat(file)
	if errors.Is(err, fs.ErrNotExist) {
		err = nil
		exists = false
		goto end
	}
	if err != nil {
		slog.Error("Stating file", "file", file, "error", err)
		goto end
	}
	exists = !info.IsDir()
end:
	return exists
}

func dirMustExist(path string) (exists bool) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = nil
		exists = false
		goto end
	}
	if err != nil {
		slog.Error("Stating dir", "path", path, "error", err)
		return false
	}
	exists = info.IsDir()
end:
	return exists
}
