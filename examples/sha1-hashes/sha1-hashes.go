// [_SHA1 散列_](http://en.wikipedia.org/wiki/SHA-1)经常用
// 生成二进制文件或者文本块的短标识。例如，[git 版本控制系统](http://git-scm.com/)
// 大量的使用 SHA1 来标识受版本控制的文件和目录。这里是 Go
// 中如何进行 SHA1 散列计算的例子。

package main

// Go 在多个 `crypto/*` 包中实现了一系列散列函数。
import "crypto/sha1"
import "fmt"

func main() {
    s := "sha1 this string"

    // 产生一个散列值得方式是 `sha1.New()`，`sha1.Write(bytes)`，
    // 然后 `sha1.Sum([]byte{})`。这里我们从一个新的散列开始。
    h := sha1.New()

    // 写入要处理的字节。如果是一个字符串，需要使用
    // `[]byte(s)` 来强制转换成字节数组。
    h.Write([]byte(s))

    // 这个用来得到最终的散列值的字符切片。`Sum` 的参数可以
    // 用来都现有的字符切片追加额外的字节切片：一般不需要要。
    bs := h.Sum(nil)

    // SHA1 值经常以 16 进制输出，例如在 git commit 中。使用
    // `%x` 来将散列结果格式化为 16 进制字符串。
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
