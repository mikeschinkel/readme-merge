#!/bin/sh

main() {
  index_filepath="$1"
  readme_dir="$2"
  do_commit="$3"

  # Run your Go application to generate README.md
  if ! /app/bin/readme-merge "${index_filepath}" "${readme_dir}" ; then
    echo "README Merge failed. README.md likely not updated." >&2
    exit 1
  fi

  if [ "${do_commit}" != "commit" ]; then
    return
  fi

  # Configure Git
  git config --global user.name 'github-action[bot]'
  git config --global user.email 'github-action@users.noreply.github.com'

  # See https://medium.com/@janloo/github-actions-detected-dubious-ownership-in-repository-at-github-workspace-how-to-fix-b9cc127d4c04
  git config --global --add safe.directory /github/workspace

  git status

  # Extract branch name from GITHUB_REF
  BRANCH_NAME=${GITHUB_REF#refs/heads/}

  # Check if README.md has changed
  if ! git diff --quiet "${readme_dir}/README.md"; then
    # Commit and push changes
    git add "${readme_dir}/README.md"
    git commit -m 'Update README.md [skip ci]'
    git remote set-url origin "https://x-access-token:${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git"
    git push origin "HEAD:${BRANCH_NAME}"
  fi
}

main "$1" "$2" "$3"