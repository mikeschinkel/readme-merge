# Layout and Syntax
Layout and syntax for README Merge has deliberately designed to be very simple. You only need:

1. Directory for Markdown files
2. Entry-point source Markdown file 
3. Other source Markdown files
4. Merge Directives in Source files

## Directory and Entry-point File:

Your source directory of Markdown files should have an entry point file which by convention README Merge names `_index.md` in the `./md` subdirectory off the repository's root. 

Note that if you use those names — e.g. `./md/_index.md` — you will not need to specify them when using the GitHub Action. 

## Other Source Markdown Files:

Other source files are just Markdown files in the same directory as the `_index.md`file that can be merged. The first special cases for how they need to be authored is discussed below. 

## Syntax for Merge Directives

Your entry point file — which I will refer to as `_index.md` from here on in this section —  should have one or more `[merge]` directives. They are the same syntax as a link but use the word `merge` as a special keyword.

### Merged Directives Start in Column One
From [the samples `_index.md` file](./samples/md/_index.md) — shown below — you can see that the `[merge]` directive **MUST** be in the first column or README Merge will ignore it. This allows you to still link the work merge to somewhere if you have that need:

#### File `./samples/md/_index.md`
```markdown
# Just a Sample README-merge Index Template

[merge](./foo.md)
[merge](./bar.md)

# Footer
```

### Merged documents may be Nested

As you can see from [the samples `foo.md` file](./samples/md/foo.md) — merged above and shown below — you can see that the merged files can also contain `[merge]` directives:

#### File `./samples/md/_foo.md`
```markdown
# Foo Template

This is the Foo template

[merge](./baz.md)
```

### Headers Must Use Hash Characters 

To avoid scope-creep and additional complexity README Merge does not support underlines to indicate headers.  Headers must use hash (`#`) characters, e.g:

```markdown
#
##
###
####
#####
```

### Headers Start in Column One 

```markdown

# Good Header

 # Bad Header
 
  # Worse Header
```

To simplify implementation headers **MUST** be placed in column one (1) or they will be ignored by README Merge. 

The reason for this is that the first non-whitespace characters in source code examples are occassionally `#` characters — which as comments in Shell scripts nd Ruby source code — so to keep parsing simple we have decided to require headers to be in column 1. This requirement also limits potential bugs.

If this becomes a problem we can potentially enhance README Merge to recognize headings with up to three (3) space characters before the (first) `#`, but we won't do that until it becomes a repeated sore point for users.

### Merged Headers Will be Demoted

One reason the author built this was he hated having to author Markdown files as parts of a whole that could not live on their own as standalone Markdown files.

Specifically, if a Markdown file with a level 1 heading — e.g. `#` — is merged in to an `_index.md` file then that level 1 heading would compete with the level 1 heading of the `_index.md` file and thus with the resultant `README.md` output file. 


To solve this, README Merge demotes any headings of a merged document by one. For example, if `usage.md` starts with `# Usage` and is merged into `_index.md` then it will appear as `## Usage` in the output `README.md`.

#### In `./md/usage.md`
```markdown
# Usage
```
#### Whereas in `./README.md`
```markdown
## Usage
```

This is further true of nested documents. This is `docker-usage.md` is merged into `usage.md` which is merged into `_index.md` and `docker-usage.md` has a header of `# Docker Usage` then in the resultant `README.md` it will appear two levels demoted, or `### Docker Usage`. 

#### In `./md/docker-usage.md`
```markdown
# Docker Usage
```
#### Whereas in `./README.md`
```markdown
## Usage
### Docker Usage
```

## Examples to Review

To understand the layout and syntax expected by README merge you can review the very simple samples files located in `./samples/md` as seen in the screenshot below:

![Tree for sample markdown source files](./md/assets/samples-source-tree.png)
_Directory tree for sample Markdown source files._


Alternately you can look at the source markdown files for this repo's `README.md` in the `./md`:

![Tree for repo's actual markdown source files](./md/assets/actual-source-tree.png)
_Directory tree for Markdown source files for this repo._


