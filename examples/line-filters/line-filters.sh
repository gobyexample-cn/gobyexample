# 试一下我们的行过滤器，首先创建一个有小写行的文件。
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# 然后使用行过滤器来得到大写的行。
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
