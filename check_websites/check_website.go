package check_websites

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	//for _, url := range urls {
	//	results[url] = wc(url)
	//}

	// CheckWebsites函数执行完成，其它goroutine还没有写入到results中
	for _, url := range urls {
		go func(u string) {
			results[url] = wc(url)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}

type result struct {
	url    string
	result bool
}

func CheckWebsites2(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// 阻塞等待数据从channel中读取
		result := <-resultChannel
		results[result.url] = result.result
	}

	return results
}
