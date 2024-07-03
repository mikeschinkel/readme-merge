#!/bin/sh

main() {
  index_filepath="$1"
  readme_dir="$2"

  # Run your Go application to generate README.md
  if ! /app/bin/readme-merge "${index_filepath}" "${readme_dir}" ; then
    echo "README Merge failed. README.md likely not updated." >&2
    exit 1
  fi

  # Configure Git
  git config --global user.name 'readme-merge-github-action[bot]'
  git config --global user.email 'readme-merge-github-action@users.noreply.github.com'

  git status

  # Check if README.md has changed
  if ! git diff --quiet "${readme_dir}/README.md"; then
    # Commit and push changes
    git add "${readme_dir}/README.md"
    git commit -m 'Update README.md [skip ci]'
    git push
  fi
}

main "$1" "$2"