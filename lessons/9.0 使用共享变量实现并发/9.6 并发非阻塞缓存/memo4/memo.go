package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}

func httpGetBody(url string) (interface{}, error) {
	log.Printf("require %s", url)
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
	for _, url := range urls {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s, %s, %d bytes \n",
			url, time.Since(start), len(value.([]byte)))
	}
}
