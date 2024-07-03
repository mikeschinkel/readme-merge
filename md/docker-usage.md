# Docker Container Usage

## Syntax
```go
docker build -t readme-merge:<your_tag> ../

docker run -v "<your_local_repo_root>:/app/<repo_in_container>" --rm \
  readme-merge:<your_tag> \
    ./<repo_in_container>/<local_markdown_source_dir>/<entry_point_doc> \
    ./<repo_in_container>/<local_dir_for_readme> \
    [commit|no_commit]
```
## Example
```go
docker build -t readme-merge:your_tag ../

docker run -v "$(pwd)/..:/app/repo" --rm \
  readme-merge:local_usage \
    ./repo/samples/md/_index.md \
    ./repo/samples/. \
    no_commit
```