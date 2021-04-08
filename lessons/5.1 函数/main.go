package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func countWordsAndImages(n *html.Node) (words, images int) {
	return 0, 0
}

func CountWordsAndImages(url string) (words, images int, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(response.Body)
	response.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func getDoc(url string) (doc *html.Node, error error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, response.Status)
	}
	doc, err = html.Parse(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return doc, nil
}

func findLinks(url string) ([]string, error) {
	doc, err := getDoc(url)
	return visit(nil, doc), err
}

func findLines(url string) {
	doc, err := getDoc(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	outline(nil, doc)
	outline2(doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func outline2(doc *html.Node) {
	forEachNode(doc, startElement, endElement)
}

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

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s> \n", depth*4, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s> \n", depth*4, "", n.Data)
	}
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	var urls = []string{"https://www.baidu.com"}
	for _, url := range urls {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findLink: %v \n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
		findLines(url)
	}
	fmt.Println(f())
	fmt.Println(f2())
	fmt.Println(f3())
}
