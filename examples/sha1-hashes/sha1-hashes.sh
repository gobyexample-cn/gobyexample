# 运行程序计算散列值并以可读 16 进制格式输出。
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94


# 你可以使用和上面相似的方式来计算其他形式的散列值。例
# 如，计算 MD5 散列，引入 `crypto/md5` 并使用 `md5.New()`
# 方法。

# 注意，如果你需要密码学上的安全散列，你需要小心的研究一下
# [哈希强度](http://en.wikipedia.org/wiki/Cryptographic_hash_function)。
