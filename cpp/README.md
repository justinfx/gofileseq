## libfileseq

#### C++ bindings for gofileseq

This is an experimental library which provides a C++ wrapper over gofileseq, using Go's buildmodes to export the support.

### Build Requirements

* [go](https://golang.org/dl/) >= 1.6.0
* [qmake](https://www.qt.io/download-open-source/) (comes with Qt) or [waf](https://waf.io/)

Go is only needed to build the library and is not a runtime dependency.

You don't have to have qmake or waf if you want to compile it with your own Makefile :-)

### Runtime Dependencies

None

### Building with qmake

```shell
$ cd <path/to/gofileseq>/cpp
$ qmake
$ make
```

### Testing with qmake

```shell
$ cd <path/to/gofileseq>/cpp/test
$ qmake
$ make
```

### Building with waf

```shell
$ cd <path/to/gofileseq>/cpp
$ waf configure --prefix=/path/to/install/to
$ waf build
$ waf test
$ waf install
```
