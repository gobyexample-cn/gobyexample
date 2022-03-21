## Go by Example 中文版

[Go by Example](https://gobyexample-cn.github.io/) 是一个通过带注释的示例程序学习 Go 语言的网站。网站包含了从简单的 Hello World 到高级特性 Goroutine、Channel 等一系列示例程序，并附带了注释说明，非常适合 Go 语言初学者。

如果您想学习 Go 语言基础知识，不要犹豫，请直接前往 [Go by Example](https://gobyexample-cn.github.io/) 开始学习！

如果您觉得本项目还不错的话，记得回来给个 Star 哦 o(*￣▽￣*)ブ

## 综述

如果你想了解 Go by Example `网站` 是如何构建的，或者想为该项目贡献代码，请查看下面的内容：

本项目包含了网站的内容和构建工具链，网站使用的是 `public` 目录下静态文件（html 等文件）的内容。它是这样被构建出来的：通过 `程序` 提取 `examples` 目录下的源码及注释，并使用 `templates` 目录下的静态文件模板将其渲染为静态文件，最终将生成的静态文件输出到 `public` 目录下。

实现此构建过程的 `程序` 位于 `tools` 目录下，构建得到的 `public` 目录下的静态文件（html 等文件），可以部署到任何支持静态内容的系统。例如 S3、CloudFront 以及任何 Web 服务器。

### 构建

[![Build Status](https://github.com/mmcgrana/gobyexample/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/mmcgrana/gobyexample/actions)

若想自行构建该网站，你需要安装 Go。然后运行下面的命令：

```console
$ tools/build
```

若想循环持续构建：

```console
$ tools/build-loop
```

在本地启动服务浏览网站

```console
$ tools/serve
```

然后再浏览器中打开 `http://127.0.0.1:8000/`

### 发布

下面的例子展示了如何将网站上传至 AWS：

```console
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ tools/upload
```

## 许可协议

该项目的著作权归 Mark McGranaghan 所有，并遵循 [CC BY-SA 3.0](http://creativecommons.org/licenses/by/3.0/) 协议。

Go Gopher 的版权归 [Renée French](http://reneefrench.blogspot.com/) 所有，并遵循 [CC BY-SA 3.0](http://creativecommons.org/licenses/by/3.0/) 协议。

## 其他语言

本项目只是 [mmcgrana](https://github.com/mmcgrana) 的 [Go by Example](https://github.com/mmcgrana/gobyexample) 项目的中文翻译。

除中文版外，该项目还有以下语言：

* [English](https://gobyexample.com) by [mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample)（原版）
* [Czech](http://gobyexamples.sweb.cz/) by [martinkunc](https://github.com/martinkunc/gobyexample-cz)
* [French](http://le-go-par-l-exemple.keiruaprod.fr) by [keirua](https://github.com/keirua/gobyexample)
* [Italian](http://gobyexample.it) by the [Go Italian community](https://github.com/golangit/gobyexample-it)
* [Japanese](http://spinute.org/go-by-example) by [spinute](https://github.com/spinute)
* [Korean](https://mingrammer.com/gobyexample/) by [mingrammer](https://github.com/mingrammer)
* [Russian](https://gobyexample.com.ru/) by [badkaktus](https://github.com/badkaktus)
* [Spanish](http://goconejemplos.com) by the [Go Mexico community](https://github.com/dabit/gobyexample)
* [Ukrainian](http://butuzov.github.io/gobyexample/) by [butuzov](https://github.com/butuzov/gobyexample)

## 致谢

感谢 [Jeremy Ashkenas](https://github.com/jashkenas) 的 [Docco](http://jashkenas.github.com/docco/)，启发了这个项目。

## 贡献说明

> 从这部分开始，后面的内容都是中文版的贡献者们给自己加的戏。
>
> 好吧，其实前面的内容也加了很多戏，没有完全根据英文版翻译。

如果你发现中文版的例子没有及时与英文版同步，或者你觉得某个例子翻译得不够好，甚至只是一个错误的文字、单词或符号，我们都 `非常欢迎` 你能够提交 pull request 以帮助我们使项目更完善，贡献流程大致如下：

1. Fork 该仓库。
1. 在 `examples` 目录下找到想要修改的例子，完成修改，这通常是以 `例子`（也就是一个目录）为单位进行修改，当然，你可以一次性修改多个例子。需要注意的是：只修改 `.go` 和 `.sh` 文件。`.hash` 文件是 `tools/build` 自动更新的，主要用于判断文件内容是否有改动；
1. 使用 `tools/build` 命令重新生成静态文件。这一步会格式化代码，并判断内容是否有改动。对于内容有改动的例子，会自动将该例子的代码提交至 `https://play.studygolang.com/` 进行测试。通过测试后，会自动更新静态文件；
1. `tools/serve` 本地预览效果；
1. 通过自测后即可提交 pull request :)

项目现由 [gobyexample-cn](https://github.com/gobyexample-cn) 维护，例子已完全与英文版同步（截止 2022-3-20），均为 79 个，可以在这里查看 [同步进度](PROGRESS.md)。

后续可能会出现与英文版同步不及时的情况，`非常欢迎` 各位同学 fork 并提交 pull request。

## 中文版的致谢

感谢本翻译项目的原作者 [everyx](https://github.com/everyx)，完成了所有文件最初的翻译，同时也感谢项目每一位 [贡献者](https://github.com/gobyexample-cn/gobyexample/graphs/contributors) 的辛勤付出。

`JetBrains` 公司为本项目提供了 free JetBrains Open Source license(s)，在此表示感谢。

[![jetbrains](jetbrains-logo/jetbrains.svg)](https://www.jetbrains.com/?from=gobyexample-cn)
[![jetbrains](jetbrains-logo/goland.svg)](https://www.jetbrains.com/?from=gobyexample-cn)
[![jetbrains](jetbrains-logo/webstorm.svg)](https://www.jetbrains.com/?from=gobyexample-cn)
