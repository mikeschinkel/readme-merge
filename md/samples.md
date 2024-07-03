# Samples

There is a very simply set of Markdown documents in the `/samples/md` directory. 

![Samples directory tree](./md/assets/samples-source-tree.png)
_Directory tree for sample Markdown source files._

## Building the sample `README.md` 
You can see how the samples work by running the following commands which will build a Docker container to run `go build` and then run `readme-merge` which will merge the sample docs into an example `README.md` file as `/samples/README.md` and then use `less` to view it:

```shell
git clone https://github.com/mikeschinkel/readme-merge
cd readme-merge/samples
./build-sample-readme.sh
```

## Requirements for running `./md/build-sample-readme.sh`
To build the sample `README.md` you need:

- To be running a flavor of *nix _(Linux, macOS, WSL, etc)_
- To have a running [Docker daemon installed](https://docs.docker.com/engine/install/).

There is not batch file or PowerShell script for use if you are running Windows but if you are motivated you can port `./samples/build-sample-readme.sh` and support a PR to include for future Windows users.  