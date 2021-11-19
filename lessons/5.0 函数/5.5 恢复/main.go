package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func main() {
	url := "https://www.bing.com"
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		fmt.Errorf("%s has type %s, not text/html", url, ct)
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}
	fmt.Println(soleTitle(doc))
}
