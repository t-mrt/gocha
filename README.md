gocha
=====

NAME
----

gocha - Generate random strings based on a pattern

DESCRIPTION
-----------

gocha makes it trivial to generate random strings.

```bash
$ gocha -n 5 "(:shimobayash2?:){5}"
:shimobayash::shimobayash2::shimobayash2::shimobayash::shimobayash2:
:shimobayash2::shimobayash2::shimobayash::shimobayash::shimobayash2:
:shimobayash2::shimobayash::shimobayash::shimobayash2::shimobayash:
:shimobayash2::shimobayash::shimobayash::shimobayash::shimobayash2:
:shimobayash::shimobayash2::shimobayash::shimobayash2::shimobayash2:
```

You can use the same regular expression as [regexp/syntax](https://golang.org/pkg/regexp/syntax/).

SYNOPSIS
--------

```
usage: gocha [<flags>] <pattern>

Flags:
      --help  Show context-sensitive help (also try --help-long and --help-man).
  -n, --number-of-lines=NUMBER-OF-LINES
              Number of lines

Args:
  <pattern>  Regular expression
```

INSTALLATION
------------

```
go get github.com/t-mrt/gocha/cmd/gocha
```

AUTHOR
------

t_ <t.us.mrt@gmail.com>
