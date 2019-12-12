## Go by Example 中文版

本项目是 [mmcgrana](https://github.com/mmcgrana) 的 [Go by Example](https://github.com/mmcgrana/gobyexample) 中文翻译。现已由[everyx](https://github.com/everyx) 完成了所有文件的初步翻译，目前正在准备重新翻译，进度条见[文件](./PROGRESS.md)，欢迎 fork 并提交 pull request。

### 贡献说明

1. 首先在· `examples` 目录下找到对应的内容，完成修改。注：只需注意 `.go` 文件和 `.sh` 文件。`.hash` 文件是在使用 `tools/build` 后自动更新的，主要用于判断文件内容是否被修改；
1. 然后使用 `tools/build` 命令重新生成网页。这一步会格式化代码，并判断是否修改，修改的内容会重新提交示例代码至 `http://play.golang.org/`，并更新静态页；
1. `tools/serve` 本地测试结果；
1. 提交 pull request :)

### License

This work is copyright Mark McGranaghan and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

The Go Gopher is copyright [Renée French](http://reneefrench.blogspot.com/) and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

### 原版传送门

- [英文版](https://gobyexample.com)
- [mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample)

### 致谢

本翻译项目原作者[everyx](https://github.com/everyx)，以及所有[贡献者](https://github.com/gobyexample-cn/gobyexample/graphs/contributors)。

---

## Go by Example 原版 README

Content and build toolchain for [Go by Example](https://gobyexample.com),
a site that teaches Go via annotated example programs.

### Overview

The Go by Example site is built by extracting code and
comments from source files in `examples` and rendering
them via the `templates` into a static `public`
directory. The programs implementing this build process
are in `tools`, along with some vendor'd dependencies
in `vendor`.

The built `public` directory can be served by any
static content system. The production site uses S3 and
CloudFront, for example.

### Building

[![Build Status](https://travis-ci.com/mmcgrana/gobyexample.svg "Travis CI status")](https://travis-ci.com/mmcgrana/gobyexample)

To build the site you'll need Go and Python installed. Run:

```console
$ go get github.com/russross/blackfriday
$ tools/build
$ open public/index.html
```

To build continuously in a loop:

```console
$ tools/build-loop
```

### Publishing

To upload the site:

```console
$ gem install aws-sdk
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ tools/upload
```

### License

This work is copyright Mark McGranaghan and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

The Go Gopher is copyright [Renée French](http://reneefrench.blogspot.com/) and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Translations

Contributor translations of the Go by Example site are available in:

* [Chinese](https://gobyexample-cn.github.io/) by [gobyexample-cn](https://github.com/gobyexample-cn)
* [Czech](http://gobyexamples.sweb.cz/) by [martinkunc](https://github.com/martinkunc/gobyexample-cz)
* [French](http://le-go-par-l-exemple.keiruaprod.fr) by [keirua](https://github.com/keirua/gobyexample)
* [Italian](http://gobyexample.it) by the [Go Italian community](https://github.com/golangit/gobyexample-it)
* [Japanese](http://spinute.org/go-by-example) by [spinute](https://github.com/spinute)
* [Korean](https://mingrammer.com/gobyexample/) by [mingrammer](https://github.com/mingrammer)
* [Spanish](http://goconejemplos.com) by the [Go Mexico community](https://github.com/dabit/gobyexample)
* [Ukrainian](http://gobyexample.com.ua/) by [butuzov](https://github.com/butuzov/gobyexample)

### Thanks

Thanks to [Jeremy Ashkenas](https://github.com/jashkenas)
for [Docco](http://jashkenas.github.com/docco/), which
inspired this project.
