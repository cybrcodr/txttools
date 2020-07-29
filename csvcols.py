#!/usr/bin/env python
#
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
"""Command line tool that prints selected column(s) of CSV file(s).

Output of selected columns are in CSV using the same delimiter. Use -f flag to
specify the columns to output. First column is 1. The order of the selected
columns is the order of the output.

Run `csvcols.py --help` for optional flags. It reads from stdin if no arguments
are provided.
"""

import csv
import getopt
import os
import sys

def show_usage():
  print 'Usage: %s <filename>' % (os.path.basename(sys.argv[0]),)
  print
  print 'Options:'
  print
  print ('\t-f, --fields <comma-separated list of column numbers>  '
         'select only these fields, first column is 1')
  print '\t-d, --delimiter delimiter character instead of ,'
  print '\t-l, --lineterm end of line terminator, windows or unix'
  print


def process(fh, idx_list, delim, lineterm):
  reader = csv.reader(fh, delimiter=delim, lineterminator=lineterm)
  writer = csv.writer(sys.stdout, delimiter=delim, lineterminator=os.linesep)
  for line_num, row in enumerate(reader):
    size = len(row)
    if size:
      out = []
      for idx in idx_list:
        if idx >= size:
          print ('\nError: line %d : column %s out of range for row size %d' %
                 (line_num + 1, idx + 1, size))
          print row
          return

        out.append(row[idx])
      if len(out) == 0 or (len(out) == 1 and out[0] is None or out[0] == ''):
        # Avoid printing out "".
        writer.writerow([])
      else:
        writer.writerow(out)


def main(argv):
  try:
    opts, args = getopt.getopt(argv[1:], 'hd:f:l:',
                               ['help', 'delimiter=', 'fields=','lineterm='])
  except getopt.GetoptError, err:
    print err
    show_usage()
    sys.exit(2)

  idx_list = []
  delim = ','
  eol = os.linesep
  for o, a in opts:
    if o in ('-h', '--help'):
      show_usage()
      sys.exit()
    elif o in ('-f', '--fields'):
      idx_list = [int(x) - 1 for x in a.split(',')]
    elif o in ('-d', '--delimiter'):
      delim = a
    elif o in ('-l', '--lineterm'):
      if a == 'windows':
        eol = '\r\n'
      elif a == 'unix':
        eol = '\n'
      else:
        print 'Invalid value for --lineterm'
        show_usage()
        sys.exit()

  argc = len(args)
  if argc == 0:
    # read from stdin
    fh = sys.stdin
    process(fh, idx_list, delim, eol)
  else:
    for filename in args:
      fh = open(filename, 'r')
      process(fh, idx_list, delim, eol)
      fh.close()


if __name__ == '__main__':
  main(sys.argv)
