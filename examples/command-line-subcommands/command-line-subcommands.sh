$ go build command-line-subcommands.go 

# 首先调用 foo 子命令。
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# 然后试一下 bar 子命令。
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# 但是 bar 不接受 foo 的 flag（enable）。
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# 接下来我们会学习程序获取参数的另一种常见方式——环境变量。