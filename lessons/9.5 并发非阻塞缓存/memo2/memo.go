package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	response, ok := memo.cache[key]
	if !ok {
		response.value, response.err = memo.f(key)
		memo.cache[key] = response
	}
	memo.mu.Unlock()
	return response.value, response.err
}

func httpGetBody(url string) (interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func main() {
	m := New(httpGetBody)
	urls := []string{
		"https://news.baidu.com/",
		"https://map.baidu.com/",
		"https://tieba.baidu.com/",
		"https://xueshu.baidu.com/",
		"https://news.baidu.com/",
		"https://map.baidu.com/",
		"https://tieba.baidu.com/",
		"https://xueshu.baidu.com/",
	}
	var n sync.WaitGroup
	for _, url := range urls {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s, %s, %d bytes \n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
