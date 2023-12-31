# json-pprint
CLI that will pretty print Json. Served as a small project to learn Golang basics.

## Installing
Make sure `$GOPATH/bin` is in your `$PATH`, otherwise just add it to your path. If `$GOPATH` is empty, you can retrieve the value like so `go env | grep GOPATH`

Run `cd json-pprint && go install`, you should now be able to use the CLI, see how to use it below.

## Usage

```
NAME:
   json-pprint - Json to pretty print

USAGE:
   json-pprint [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --write, -w             write result to file instead of stdout, by default overwrites source file (default: false)
   --spaces, -s            use 2 spaces instead of tab (default: false)
   --file value, -f value  file to read JSON from
   --out value, -o value   output file for result
   --help, -h              show help
```