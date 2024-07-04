package main

import (
	"fmt"
	"log/slog"
	"os"
)

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
	slog.Info("Saving", "readme_file", file.Filepath())
	err = file.Save(readme)
	if err != nil {
		slog.Error("Saving README.md", "readme_file", readmeDir, "error", err)
		os.Exit(3)
	}
	slog.Info("Done")
}

func parseArgs() (indexFile, readmePath string) {
	var msg string
	switch {
	case len(os.Args) < 3:
		msg = "Not enough arguments"
		goto err
	case !fileMustExist(os.Args[1]):
		msg = fmt.Sprintf("File '%s' does not exist", os.Args[1])
		goto err
	case !dirMustExist(os.Args[2]):
		msg = fmt.Sprintf("Directory '%s' does not exist", os.Args[2])
		goto err
	}
	indexFile = os.Args[1]
	readmePath = os.Args[2]
	goto end
err:
	_, _ = fmt.Fprintf(os.Stderr, "Usage: readme-merge <index_file> <readme_path>\n\n\t%s\n", msg)
	os.Exit(1)
end:
	return indexFile, readmePath
}
