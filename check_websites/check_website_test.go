package check_websites

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "a.b.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"a.b.com",
		"2.3.com",
		"3.4.com",
	}

	actualResult := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResult)
	if got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}

	expectedResult := map[string]bool{
		"a.b.com": false,
		"2.3.com": true,
		"3.4.com": true,
	}

	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Fatalf("Wanted %v, got %v", expectedResult, actualResult)
	}
}

func slowMockWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// 基准测试命令：go test check_website_test.go check_website.go -bench=.
// 在运行基准测试时：fatal error: concurrent map writes   发生了同时多个goroutine写入map的情况
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "111"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites2(slowMockWebsiteChecker, urls)
	}
}
