package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan<- result
}

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	requests chan request
}

type Func func(key string) (interface{}, error)

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}

}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
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
