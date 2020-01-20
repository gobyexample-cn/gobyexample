# 生成的程序返回的输出，和我们直接通过命令行运行这些程序的输出是相同的。
$ go run spawning-processes.go 
> date
Wed Oct 10 09:53:11 PDT 2012

> grep hello
hello grep

> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go
