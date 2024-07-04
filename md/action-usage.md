# GitHub Action Usage

In order to configure this GitHub Action to work for your repository you need to do the following three (3) things:

1. Add a directory of source Markdown files and an entry point file
1. Set Read/Write Access
2. Add a workflow .YAML file

## 1. Create Source Dir and Entry-point File
README Merge reads an entry point file in a subdirectory that defaults — if you do not specify otherwise in your workflow — to `./md/_index.md` where `/md` is a subdirectory off the root.

For more details about the entry-point file and document syntax see [Layout and Syntax](./layout-syntax.md).

## 2. Set Read/Write Access
As the entire purpose of this action is to build the `README.md` this action cannot perform its purpose unless it can update the GitHub repo. **If you have security concerns about write access see the [security section below](#security-considerations).**

To configure your GitHub repo for to allow writes, visit your repo's Settings from the top GitHub menu, select GitHub Actions and then follow the instructions in this screenshot:

![In your repo's Settings select GitHub Actions > General, select the Read and write permissions radio button in the Workflow Permissions section, and click the Save button.](./assets/read-write-access-setting.png)

## 3. Workflow using the Github Action
Save the following code as `.github/workflows/generate-readme.yaml` in your repo and commit.  

No changes necessary if:
1. Your source Markdown files are in a subdirectory `/md` off your repo's root,
2. Your entry point document in your subdirectory is named `_index.md`,
3. You use `main` as your primary branch, and 
4. You want to update the `README.md` on push the `main`.

If you have more complex needs and/or use-cases then modifications to this workflow will obviously be required:

```go
name: Generate README

on:
  push:
    branches:
    - main
    - test-actions

jobs:
  generate-readme:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout repository
      uses: actions/checkout@v4

    - name: Generate README
      uses: mikeschinkel/readme-merge@main
      with:
        index_filepath: './md/_index.md'
        readme_dir: '.'
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

## SECURITY CONSIDERATIONS
This GitHub Action requires read and **write** access to your repository so it can update the `README.md` file. The updates happen in the shell script `/bin/entrypoint.sh` for your security review.

If you cannot set read/write access because your security team does not allow it, or you simply do not want to allow write access to GitHub Actions on your repo _then do not use this GitHub Action_ and instead either use the Docker container or the `readme-merge` executable elsewhere in your CI/CD and/or README authoring process. 

