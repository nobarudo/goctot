# goctot

goctot displays CSV files in a table format.

## Install

```bash
go install github.com/nobarudo/goctot@latest
```

If the command is not found, make sure `$HOME/go/bin` is in your PATH.

## Usage

```bash
goctot data.csv
goctot data.csv -m
goctot data.csv -m -o output.md
cat data.csv | goctot
```

```
goctot [file] [flags]
goctot [command]
```

### Flags

```
  -h, --help            help for goctot
  -m, --markdown        format to Markdown.
  -n, --no-header       no header option.
  -o, --output string   Specify the file to output from standard output.
```

### Subcommand

```
  help        Help about any command
  version     display version
```
