# readme-merge
CLI tool to allow merging child documents from a subdirectory into one `README.md` in the root of your repository.

**BE CAREFUL**: This will _**overwrite**_ your `README.md` file. 

## Usage
### Syntax
```go
readme-merge <index_file> <readme_path>
```
### Example
```go
readme-merge md/_index.md .
```

## Where Used
- [github.com/mikeschinkel/php-file-scoped-visibility-rfc/](https://github.com/mikeschinkel/php-file-scoped-visibility-rfc/)

## License 
MIT
