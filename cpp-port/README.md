## libfileseq

#### C++ port of gofileseq

[Latest API Docs](http://justinfx.com/gofileseq/cpp)

### Build Requirements

* [waf](https://waf.io)
* [Google Test](https://github.com/google/googletest) - If you want to run the unittests

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
