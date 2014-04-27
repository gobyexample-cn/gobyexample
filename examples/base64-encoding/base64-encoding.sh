# 标准 base64 编码和 URL 兼容 base64 编码的编码字符串存在
# 稍许不同（后缀为 `+` 和 `-`），但是两者都可以正确解码为
# 原始字符串。
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
