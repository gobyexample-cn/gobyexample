// [_SHA256 散列（hash）_](https://en.wikipedia.org/wiki/SHA-2)
// 经常用于生成二进制文件或者文本块的短标识。
// 例如，TLS/SSL 证书使用 SHA256 来计算一个证书的签名。
// 这是 Go 中如何进行 SHA256 散列计算的例子。

package main

// Go 在多个 `crypto/*` 包中实现了一系列散列函数。
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// 这里我们从一个新的散列开始。
	h := sha256.New()

	// 写入要处理的字节。如果是一个字符串，
	// 需要使用 `[]byte(s)` 将其强制转换成字节数组。
	h.Write([]byte(s))

	// `Sum` 得到最终的散列值的字符切片。`Sum` 接收一个参数，
	// 可以用来给现有的字符切片追加额外的字节切片：但是一般都不需要这样做。
	bs := h.Sum(nil)

	// SHA256 值经常以 16 进制输出，例如在 git commit 中。
	// 我们这里也使用 `%x` 来将散列结果格式化为 16 进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
