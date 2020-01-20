# 测试这个程序前，最好将这个程序编译成二进制文件，然后再运行这个程序。
$ go build command-line-flags.go

# 首先以给所有标志赋值的方式，尝试运行构建的程序。
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# 注意，如果你省略一个标志，那么这个标志的值自动的设定为他的默认值。
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# 尾随的位置参数可以出现在任何标志后面。
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# 注意，`flag` 包需要所有的标志出现位置参数之前（否则，这个标志将会被解析为位置参数）。
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# 使用 `-h` 或者 `--help` 标志来得到自动生成的这个命令行程序的帮助文本。
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# 如果你提供了一个没有使用 `flag` 包声明的标志，
# 程序会输出一个错误信息，并再次显示帮助文本。
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
