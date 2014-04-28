# 运行这端文件写入代码。
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# 然后检查写入文件的内容。
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# 下面我们将看一些文件 I/O 的想法，就像我们已经看过的 
# `stdin` 和 `stdout` 流。
