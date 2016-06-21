seqinfo
========

A command-line utility, using [gofileseq](https://github.com/justinfx/gofileseq), to parse one or more sequence pattern strings and list plain text or json info.

Example sequence formats
---------------------------------

* /path/to/some/file_foo.0100.exr
* /path/to/some/file_foo.1-100#.jpg
* /path/to/some/file_foo.1-100@@@.tif
* relative_file.1-10,50-100x2#.tif
* file_without_frame.ext

Usage
---------

```bash
# Single string pattern, in plain-text
$ seqinfo "/path/to/seq.1-100#.exr"
Source = /path/to/seq.1-100#.exr
    String = /path/to/seq.1-100#.exr
    Dirname = /path/to/
    Basename = seq.
    Range = 1-100
    Padding = #
    Ext = .exr
    Start = 1
    End = 100
    Len = 100
    ZFill = 4
    HasRange = true

# Multiple patterns, in json format
$ seqinfo --json "/path/to/file.exr" "/path/to/seq.-10-200x2@@.exr"
{
    "/path/to/file.exr": {
        "error": "",
        "string": "/path/to/file.exr",
        "dir": "/path/to/",
        "base": "file",
        "range": "",
        "pad": "",
        "ext": ".exr",
        "start": 0,
        "end": 0,
        "length": 1,
        "zfill": 0,
        "hasRange": false
    },
    "/path/to/seq.-10-200x2@@.exr": {
        "error": "",
        "string": "/path/to/seq.-10-200x2@@.exr",
        "dir": "/path/to/",
        "base": "seq.",
        "range": "-10-200x2",
        "pad": "@@",
        "ext": ".exr",
        "start": -10,
        "end": 200,
        "length": 106,
        "zfill": 2,
        "hasRange": true
    }
}
```
