# rowman
[![build](https://github.com/kobayashi-hash/rowman/actions/workflows/build.yaml/badge.svg)](https://github.com/kobayashi-hash/rowman/actions/workflows/build.yaml)

[![pages-build-deployment](https://github.com/kobayashi-hash/rowman/actions/workflows/pages/pages-build-deployment/badge.svg)](https://github.com/kobayashi-hash/rowman/actions/workflows/pages/pages-build-deployment)

[![Coverage Status](https://coveralls.io/repos/github/kobayashi-hash/rowman/badge.svg)](https://coveralls.io/github/kobayashi-hash/rowman)

[![update version](https://github.com/kobayashi-hash/rowman/actions/workflows/update-version.yaml/badge.svg)](https://github.com/kobayashi-hash/rowman/actions/workflows/update-version.yaml)

A CLI tool to process and filter CSV data.

## Overview
This tool reads a CSV file and outputs a specified number of rows or columns. It can also filter and display rows that contain a specific string.
## Usage
```
rowman [OPTION] <csv_file>
-r, --rows <N>      output the first N rows.
-c, --cols <N>      output the first N columns.
-f, --filter <TEXT> output only rows that contain the specified TEXT.
-h, --help          print this message.
```
