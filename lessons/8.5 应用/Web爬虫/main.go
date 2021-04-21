package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func Extract(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, response.Status)
	}
	doc, err := html.Parse(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := response.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil, 0)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node), depth int) {
	if depth > 5 {
		return
	}
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, depth+1)
	}
	if post != nil {
		post(n)
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := Extract(url)
	<-tokens
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- []string{"https://www.baidu.com", "https://baidu.com"} }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
