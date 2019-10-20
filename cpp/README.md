## libfileseq

#### C++ port of gofileseq

[Latest API Docs](http://justinfx.com/gofileseq/cpp)

### Build Requirements

* [waf](https://waf.io) or [cmake](https://cmake.org/)

### Runtime Dependencies

None

### Building with waf

```shell
$ cd <path/to/gofileseq>/cpp
$ waf configure --prefix=/path/to/install/to
$ waf build
$ waf test
$ waf install
```

### Building with cmake

```shell
$ cd <path/to/gofileseq>/cpp
$ mkdir build && cd build
$ cmake ..
$ make
$ ./test/fileseq_test
```
