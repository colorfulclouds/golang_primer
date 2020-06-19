// search
//文件夹search下的每个代码文件都使用search作为包名
package search

import (
	"log"
	"sync"
)

//包级变量
//Matcher类型的映射（map） key：string value：Matcher | map[key]value
//大写字母开头 公开
//小写字母开头 不公开（可以间接访问）
var matchers = make(map[string]Matcher)

//searchTerm 参数 string 参数类型

func Run(searchTerm string) {
	//下面的函数返回两个值 Feed类型的切片 Feed：引用类型
	//声明变量并赋予初始值 与var声明的变量没有区别
	// := 简化变量声明运算符
	feeds, err := RetrieveFeeds() //search中的RetrieveFeeds函数

	if err != nil {
		log.Fatal(err)
	}
	/*
		根据经验，如果需要声明初始值为零值	的变量，应该使用 var 关键字声明变量；
		如果提供确切的非零值初始化变量或者使用函数返回值创建变量，应该使用简化变量声明运算符
	*/
	//chan(通道类型) map(映射) slice(切片) 都是引用类型
	//chan 自身实现一组带类型的值 用于在goroutine之间传递数据
	results := make(chan *Results)

	var waitGroup sync.WaitGroup //WaitGroup是一个计数信号量 可以统计所有goroutine是不是都完成了工作

	waitGroup.Add(len(feeds)) //WaitGroup的值设置为将要启动的goroutine的数量

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"] //声明在上面
		}

		go func(matcher, Matcher, feed *Feed) {
			Match(macher, feed, searchTerm, results)
			waitGroup.Done() //没有传入 waitGroup 但是依然可以访问这个值
			/*
				因为有了闭包，函数可以直接访问到那些没有作为参数传入
				的变量。匿名函数并没有拿到这些变量的副本，而是直接访问外层函数作用域中声明的这些变量
				本身。
			*/

			/*
				因为matcher和feed变量每次调用时值不相同，所以并没有使用闭包形式访问这两个变量
			*/
		}(matcher, feed) //go中变量传递为值传递
	}

	go func() {
		//该匿名函数使用闭包访问了WaitGroup和results变量
		waitGroup.Wait() //等待所有的任务完成 会导致goroutine阻塞（暂停向下执行），直到WaitGroup内部的计数为0
		close(results)   //关闭通道的方式，通知Display函数 close为goroutine内置函数
	}()

	Display(results)

}

func main() {
	fmt.Println("Hello World!")

}
