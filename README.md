# shini -- A minimalistic TOML processor

shini is a tool for reading [TOML/INI](https://en.wikipedia.org/wiki/TOML)
configuration files from command-line.

**Example**

```
$ cat >config.ini <<EOT
[servers.alpha]
    ip   = "192.168.0.5"
    name = "fnord"

[servers.beta]
    ip   = "192.168.0.23"
    name = "fnord"

EOT

$ shini -f config.ini servers.alpha.ip || echo "key not found"
192.168.0.5
```
**Usage**

```
shini [-f <file>] [<keys...>]
shini [-f <file>] -q <query>

If '-f' is omitted, shini reads from stdin.
If <keys> is omitted, shini dumps all keys and values.
If a requested key does not exist, shini prints a newline and
returns exit code 1.
```

## Queries

shini uses the [go-toml](https://github.com/pelletier/go-toml) library
by Thomas Pelletier and therefore also supports queries, using `-q` option.

The syntax of a query begins with a root token, followed by any number
of sub-expressions:

```
$
                 Root of the TOML tree.  This must always come first.
.name
                 Selects child of this node, where 'name' is a TOML key
                 name.
['name']
                 Selects child of this node, where 'name' is a string
                 containing a TOML key name.
[index]
                 Selcts child array element at 'index'.
..expr
                 Recursively selects all children, filtered by an a union,
                 index, or slice expression.
..*
                 Recursive selection of all nodes at this point in the
                 tree.
.*
                 Selects all children of the current node.
[expr,expr]
                 Union operator - a logical 'or' grouping of two or more
                 sub-expressions: index, key name, or filter.
[start:end:step]
                 Slice operator - selects array elements from start to
                 end-1, at the given step.  All three arguments are
                 optional.
```

**Example**

```
$ cat config.ini | shini -q '$..[ip,name]'
192.168.0.5
fnord
192.168.0.23
fnord
```

For more details on those queries, have a look at the
[go-toml/query documentaion](https://godoc.org/github.com/pelletier/go-toml/query)


## Install

Assuming you have Go installed, you can install shini by using
the `go get` method:

        go get -u github.com/5nord/shini


## Wishlist

This tool was inspired by [jq](https://stedolan.github.io/jq/) and my reluctance
to write a TOML-parser in shell-script. Although it serves its purpose, there
are some features, which might come in handy someday:

 - [ ] Write support for configuration values
 - [ ] Shell-friendly formatting of array values
 - [ ] Additional output formats (JSON, CSV, ...)
