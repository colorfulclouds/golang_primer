package search

//下面的是一个空结构
//空结构创建实例时 不会分配任何内存 这种结构很适合创建没有任何状态的类型
type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
