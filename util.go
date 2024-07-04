package main

import (
	"errors"
	"io"
	"io/fs"
	"log/slog"
	"os"
)

func mustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		slog.Warn("Error occurred when calling mustClose()", "error", err)
	}
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