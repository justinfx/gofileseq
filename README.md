Fileseq
=======

A Go language library for parsing file sequence strings commonly
used in VFX and animation applications.

Ported from the original fileseq Python library:
https://github.com/sqlboy/fileseq

Status
======

Work in progress

Frame Range Shorthand
=====================

Support for:

* Standard: 1-10
* Comma Delimted: 1-10,10-20
* Chunked: 1-100x5
* Filled: 1-100y5
* Staggered: 1-100:3 (1-100x3, 1-100x2, 1-100)
* Negative frame numbers: -10-100
* Padding: #=4 padded, @=single pad

Documentation
=============

http://godoc.org/github.com/justinfx/gofileseq
