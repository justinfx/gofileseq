Fileseq  ![Travis Build Status](https://api.travis-ci.org/justinfx/gofileseq.svg)
=======

A Go language library for parsing file sequence strings commonly
used in VFX and animation applications.

Ported from the original fileseq Python library:
https://github.com/sqlboy/fileseq

Documentation
-------------

http://godoc.org/github.com/justinfx/gofileseq

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

[ ![Download](https://api.bintray.com/packages/justinfx/utils/seqls/images/download.svg) ](https://bintray.com/justinfx/utils/seqls/_latestVersion)

```bash
go get github.com/justinfx/gofileseq/cmd/seqls
seql some/images
# some/images/file_foo.1-100@@@.tif
```
