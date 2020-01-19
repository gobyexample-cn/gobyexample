# 运行程序计算散列值，并以可读的 16 进制格式输出。
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94


# 你可以使用和上面相似的方式来计算其他形式的散列值。
# 例如，计算 MD5 散列，引入 `crypto/md5` 并使用 `md5.New()` 方法。

# 注意，如果你需要密码学上的安全散列，你需要仔细的研究一下
# [加密散列函数](http://en.wikipedia.org/wiki/Cryptographic_hash_function)。
