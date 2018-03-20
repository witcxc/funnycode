package main

import (
	"fmt"
	"golang.org/x/net/html"
	_ "log"
	"net/http"
	"os"
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
func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			fmt.Printf("%+v", n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
}
