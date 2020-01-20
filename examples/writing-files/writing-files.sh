# 运行这段文件写入代码。
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

# 我们刚刚看到了文件 I/O 思想，
# 接下来，我们看看它在 `stdin` 和 `stdout` 流中的应用。