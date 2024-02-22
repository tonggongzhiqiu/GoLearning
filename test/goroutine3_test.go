package test

import (
	"reflect"
	"testing"
	"time"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// 直接使用会造成竞争写入
		//go func(u string) {
		//	results[u] = wc(u)
		//}(url)
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)

	}
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	time.Sleep(2 * time.Second)
	return results
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://baidu.com",
		"http://baidubaidu.com",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)
	if want != got {
		t.Fatalf("wanted %v, got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://baidu.com":        true,
		"http://baidubaidu.com":   true,
		"waat://furhurterwe.geds": false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
