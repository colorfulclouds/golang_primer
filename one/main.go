// main
package main

/*
main函数保存在名为main的包里 如果不在main包 构建工程不会生成可执行文件

包类似namespace
不同包中可以相同的名字标识符 利用包名区分
*/

/*
这个技术是为了让 Go 语言对包做初始化操作，但是并不使用包里的标识符。为了让程序的
可读性更强，Go 编译器不允许声明导入某个包却不使用。下划线让编译器接受这类导入，并且
调用对应包内的所有代码文件里定义的 init 函数。对这个程序来说，这样做的目的是调用
matchers 包中的 rss.go 代码文件里的 init 函数，注册 RSS 匹配器，以便后用。我们后面会展
示具体的工作方式。
*/

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

//每个代码文件中的init函数都会在main函数执行前调用
func init() {
	log.SetOutput(os.Stdout) //日志到标准输出
}

func main() {
	// fmt.Println("Hello World!")
	search.Run("president")
}
