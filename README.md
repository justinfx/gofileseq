Fileseq  
=======
[![Go project version](https://badge.fury.io/go/github.com%2Fjustinfx%2Fgofileseq.svg)](https://badge.fury.io/go/github.com%2Fjustinfx%2Fgofileseq) 
[![GoDoc](https://godoc.org/github.com/justinfx/gofileseq?status.svg)](https://godoc.org/github.com/justinfx/gofileseq) 
[![Travis Build Status](https://api.travis-ci.org/justinfx/gofileseq.svg)](https://www.travis-ci.org/justinfx/gofileseq) 
[![Go Report](https://goreportcard.com/badge/github.com/justinfx/gofileseq)](https://goreportcard.com/report/github.com/justinfx/gofileseq)


A Go language library for parsing file sequence strings commonly
used in VFX and animation applications.

Ported from the original fileseq Python library:
https://github.com/sqlboy/fileseq

For C++, see [C++ Support](#c-support)

Frame Range Shorthand
---------------------

Support for:

* Standard: 1-10
* Comma Delimted: 1-10,10-20
* Chunked: 1-100x5
* Filled: 1-100y5
* Staggered: 1-100:3 (1-100x3, 1-100x2, 1-100)
* Negative frame numbers: -10-100
* Padding: #=4 padded, @=single pad

Sequence Formats
----------------

Sequences of files are expected to follow a pattern similar to:

* /path/to/some/file_foo.0100.exr
* /path/to/some/file_foo.1-100#.jpg
* /path/to/some/file_foo.1-100@@@.tif

Install
-------

```bash
go get github.com/justinfx/gofileseq
```

Also included is a [`seqls`](https://github.com/justinfx/gofileseq/tree/master/cmd/seqls) tool, which uses `gofileseq` to produce an executable, used for listing sequences on the filesystem:

And [`seqinfo`](https://github.com/justinfx/gofileseq/tree/master/cmd/seqinfo) tool, which can print plain-text or json parsed information about one or more sequence strings.

Download the latest [binary release](https://github.com/justinfx/gofileseq/releases/latest)

Or install from source:
```bash
go get github.com/justinfx/gofileseq/cmd/seqls
seql some/images
# some/images/file_foo.1-100@@@.tif
```

C++ Support
------------

A C++ pure port of gofileseq is also available in the [cpp-port](https://github.com/justinfx/gofileseq/tree/master/cpp-port) dir
