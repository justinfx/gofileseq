## libfileseq

#### C++ bindings for gofileseq

This is an experimental library which provides a C++ wrapper over gofileseq, using Go's buildmodes to export the support.

### Build Requirements

* go >= 1.6.0
* qmake (comes with Qt) to generate the Makefile

Go is only needed to build the library and is not a runtime dependency.

You don't have to have qmake if you want to compile it with your own Makefile :-)

### Runtime Dependencies

None

### Building

```shell
$ cd <path/to/gofileseq>/cpp
$ qmake
$ make
```

### Testing

```shell
$ cd <path/to/gofileseq>/cpp/test
$ qmake
$ make
```
