package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const brandKeyword = "merge"

var codeMatcher = regexp.MustCompile("^```")
var headerMatcher = regexp.MustCompile(`^ {0,3}#`)
var linkMatcher = regexp.MustCompile(`^( {0,3}!?)\[([^]]*?)]\(([^)]+?)\)`)

func (f File) Parse() (lines string, err error) {
	var scanner *bufio.Scanner
	var line string
	var file *os.File
	var inCode bool
	var next bool

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
		if codeMatcher.MatchString(line) {
			inCode = !inCode
		}
		if inCode {
			f.Builder.WriteString(line)
			f.Builder.WriteByte('\n')
			continue
		}
		matches := linkMatcher.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch {
			case len(match) == 0:
				continue
			case match[2] == brandKeyword && strings.TrimSpace(match[1]) != "!":
				file := filepath.Join(f.Dir, match[3])
				nf := NewInputFile(file, f.Level+1)
				lines, err = nf.Parse()
				f.Builder.WriteString(lines)
				next = true
				break
			default:
				replacement := fmt.Sprintf("%s[%s](%s)", match[1], match[2], filepath.Join(f.Dir, match[3]))
				line = strings.Replace(line, match[0], replacement, -1)
			}
		}
		if next {
			next = false
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
