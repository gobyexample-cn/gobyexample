# 运行程序计算散列值，并以可读的 16 进制格式输出。
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...


# 你可以使用和上面相似的方式来计算其他形式的散列值。
# 例如，计算 SHA512 散列，
# 引入 `crypto/sha512` 并使用 `sha512.New()` 方法。


# 注意，如果你需要密码学上的安全散列，你需要仔细的研究一下
# [加密散列函数](https://en.wikipedia.org/wiki/Cryptographic_hash_function)。
