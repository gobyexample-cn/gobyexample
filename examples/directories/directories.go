// 对于操作文件系统中的 *目录* ，Go 提供了几个非常有用的函数。

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 在当前工作目录下，创建一个子目录。
	err := os.Mkdir("subdir", 0755)
	check(err)

	// 创建这个临时目录后，一个好习惯是：使用 `defer` 删除这个目录。
	// `os.RemoveAll` 会删除整个目录（类似于 `rm -rf`）。
	defer os.RemoveAll("subdir")

	// 一个用于创建临时文件的帮助函数。
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// 我们还可以创建一个有层级的目录，使用 `MkdirAll` 函数，并包含其父目录。
	// 这个类似于命令 `mkdir -p`。
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` 列出目录的内容，返回一个 `os.DirEntry` 类型的切片对象。
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` 可以修改当前工作目录，类似于 `cd`。
	err = os.Chdir("subdir/parent/child")
	check(err)

	// 当我们列出 *当前* 目录，就可以看到 `subdir/parent/child` 的内容了。
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `cd` 回到最开始的地方。
	err = os.Chdir("../../..")
	check(err)

	// 当然，我们也可以遍历一个目录及其所有子目录。
	// `Walk` 接受一个路径和回调函数，用于处理访问到的每个目录和文件。
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// `filepath.Walk` 遍历访问到每一个目录和文件后，都会调用 `visit`。
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
