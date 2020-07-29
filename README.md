# txttools

This repository contains the following command line tools for processing text
files.

This is not an officially supported Google product.

## csvcols.py

This tool can be used for selecting column(s) from CSV file(s) or stdin. It is
similar to the Unix cut command but a field is represented by a CSV column.

It was written back in Python 2.4 and still works in Python 2.7. It does not
work in Python 3. This is the tool that I've used most over the years doing
devops type of work. I plan to develop a Go version of this.

Run `csvcols.py --help` for usage. If a filename is not specified, it will read
from stdin.

```sh
$ cat commas.csv
2015,New York University,"75 3rd Ave, New York, NY 10003"
2017,Pomona College,"170 E 6th Street #47, Claremont, CA 91711"
2018,Barnard College,"5085 Altschul, New York, NY 10027"

$ csvcols.py -f2,3 commas.csv
New York University,"75 3rd Ave, New York, NY 10003"
Pomona College,"170 E 6th Street #47, Claremont, CA 91711"
Barnard College,"5085 Altschul, New York, NY 10027"

$ cat hash.csv
2015#New York University#75 3rd Ave, New York, NY 10003
2017#Pomona College#"170 E 6th Street #206, Claremont, CA 91711"
2018#Barnard College#5085 Altschul, New York, NY 10027

$ cat hash.csv | csvcols.py -f1,3 -d#
2015#75 3rd Ave, New York, NY 10003
2017#"170 E 6th Street #206, Claremont, CA 91711"
2018#5085 Altschul, New York, NY 10027
```

## lset

This tool provides set operations between 2 files, where each input file is a
set of line items. If a file has duplicated line items, those will be treated as
the same value. Resulting values are printed to stdout where each value is a
line.

Run `lset -help` for list of supported operations.

```sh
$ lset --help
Usage: lset <command> <file1> <file2>

where <command> is one of:
        diff  - shows unique lines that are different between file1 (-) and file2 (+)
        minus - shows unique lines in file1 that are not in file2
        cross - shows unique lines that are common between file1 and file2
```

## ljoin

This tool joins lines in a given file with a separator into a single line
output. This is handy for constructing a list of arguments from a file to be
passed on to another command.

```sh
$ cat file
A.txt
B.txt
$ echo `ljoin -s=, file`
A.txt,B.txt
```

## lrand

The lrand command line tool selects up to given n number of line items from a
given list of line items either in a file or from stdin.
