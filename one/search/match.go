package search

import (
	"log"
)

type Result struct {
	Field   string
	Content string
}

//这个方法输入一个指向 Feed
//类型值的指针和一个 string 类型的搜索项。
//这个方法返回两个值：一个指向 Result 类型值
//的指针的切片，另一个是错误值。
/*命名接口的时候，也需要遵守 Go 语言的命名惯例。
如果接口类型只包含一个方法，那么这
个类型的名字以 er 结尾。*/
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)

	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
	
	func Display(results chan *Result){
		for result := range results{
			fmt.Printf("%s:\n%s\n\n", result.Field , result.Content)
		}
	}
}
