// 在前面的示例中，我们研究了配置简单的 [HTTP 服务器](http-servers)。
// HTTP 服务器对于演示 `context.Context` 的用法很有用的，
// `context.Context` 被用于控制 cancel。
// `Context` 跨 API 边界和协程携带了：deadline、取消信号以及其他请求范围的值。

package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// `net/http` 机制为每个请求创建了一个 `context.Context`，
	// 并且可以通过 `Context()` 方法获取并使用它。
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// 等待几秒钟，然后再将回复发送给客户端。
	// 这可以模拟服务器正在执行的某些工作。
	// 在工作时，请密切关注 context 的 `Done()` 通道的信号，
	// 一旦收到该信号，表明我们应该取消工作并尽快返回。
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// context 的 `Err()` 方法返回一个错误，
		// 该错误说明了 `Done` 通道关闭的原因。
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// 跟前面一样，我们在 `/hello` 路由上注册 handler，然后开始提供服务。
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
