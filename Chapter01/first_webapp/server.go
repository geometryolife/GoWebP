// 声明程序所属的包
package main

// 导入程序所需的包
// 既可以是标准库的包，也可以是第三方库的包
// fmt -> 让程序可以使用 Fprintf 等函数对 I/O 进行格式化
// net/http -> 让程序可以与 HTTP 进行交互
import (
	"fmt"
	"net/http"
)

// 定义处理器函数
// writer 参数 -> 接口
// request 参数 -> 指向 Request 结构的指针
// handler 函数：
// 从 Request 结构中提取相关的信息，创建一个 HTTP 响应
// 通过 ResponseWriter 接口将响应返回给客户端
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	// 将 handler 函数设置成根（root）URL（/）被访问时的处理器
	http.HandleFunc("/", handler)
	// 启动服务器并让它监听系统的 8080 端口
	http.ListenAndServe(":8080", nil)
}
