package readme_merge

import (
	"bufio"
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

type Merger interface {
	MergeWithLevel(level int) (string, error)
	Reader() (io.Reader, error)
	Builder() *strings.Builder
	GetChild(string) (Merger, error)
	CloseReader()
}

func Merge(m Merger, dir string, level int) (lines string, err error) {
	var line string
	var inCode bool
	var next bool
	var scanner *bufio.Scanner

	reader, err := m.Reader()
	if err != nil {
		err = fmt.Errorf("getting reader '%s': %w", line, err)
		goto end
	}
	defer m.CloseReader()
	scanner = bufio.NewScanner(reader)

	// Debug: Check if the scanner is correctly initialized
	if scanner == nil {
		err = fmt.Errorf("scanner not initialized")
		goto end
	}
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			m.Builder().WriteByte('\n')
			continue
		}
		if codeMatcher.MatchString(line) {
			inCode = !inCode
		}
		if inCode {
			m.Builder().WriteString(line)
			m.Builder().WriteByte('\n')
			continue
		}
		matches := linkMatcher.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch {
			case len(match) == 0:
				continue
			case match[2] == brandKeyword && strings.TrimSpace(match[1]) != "!":
				var child Merger
				key := filepath.Join(dir, match[3])
				child, err = m.GetChild(key)
				if err != nil {
					err = fmt.Errorf("getting child '%s': %w", line, err)
					goto end
				}
				lines, err = child.MergeWithLevel(level + 1)
				if err != nil {
					err = fmt.Errorf("merging line '%s': %w", line, err)
					goto end
				}
				m.Builder().WriteString(lines)
				next = true
				break
			default:
				replacement := fmt.Sprintf("%s[%s](%s)", match[1], match[2], filepath.Join(dir, match[3]))
				line = strings.Replace(line, match[0], replacement, -1)
			}
		}
		if next {
			next = false
			continue
		}
		if level != 0 && headerMatcher.MatchString(line) {
			line = strings.Repeat("#", level) + line
		}
		m.Builder().WriteString(line)
		m.Builder().WriteByte('\n')
	}

	if err = scanner.Err(); err != nil {
		err = fmt.Errorf("scanning index file; %w", err)
		goto end
	}

	lines = m.Builder().String()
end:
	return lines, err
}
