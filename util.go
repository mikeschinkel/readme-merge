package readme_merge

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
)

func fileExists(file string) (exists bool, err error) {
	info, err := os.Stat(file)
	if errors.Is(err, fs.ErrNotExist) {
		err = nil
		exists = false
		goto end
	}
	if err != nil {
		err = fmt.Errorf("stating file '%s'; %w", file, err)
		goto end
	}
	exists = !info.IsDir()
end:
	return exists, err
}

func FileMustExist(file string) (exists bool) {
	exists, err := fileExists(file)
	if err != nil {
		slog.Error("Stating file", "file", file, "error", err)
		goto end
	}
end:
	return exists
}

func DirMustExist(path string) (exists bool) {
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
