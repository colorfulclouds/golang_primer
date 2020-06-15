// search
//文件夹search下的每个代码文件都使用search作为包名
package search

import (
	"log"
	"sync"
)

//包级变量
//Matcher类型的映射（map） key：string value：Matcher
//大写字母开头 公开
//小写字母开头 不公开（可以间接访问）
var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Results)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		macher , exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		
		go func(matcher , Matcher , feed *Feed) {
			Match(macher , feed , searchTerm , results)
			waitGroup.Done()
		}(matcher , feed)
	}
	
	go func() {
		waitGroup.Wait(
			close(results)
		)
	}()
	
	Display(results)

}

func main() {
	fmt.Println("Hello World!")
}
