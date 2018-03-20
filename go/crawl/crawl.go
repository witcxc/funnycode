package main

import (
	_ "bytes"
	"fmt"
	"github.com/saintfish/chardet"
	"github.com/witcxc/mahonia"
	"golang.org/x/net/html"
	_ "golang.org/x/text/encoding/simplifiedchinese"
	_ "golang.org/x/text/encoding/traditionalchinese"
	_ "golang.org/x/text/transform"
	_ "io/ioutil"
	_ "log"
	"net/http"
	"os"
	"strings"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node) string) string {
	title := ""
	if pre != nil {
		tmp := pre(n)
		if tmp != "" {
			title = tmp
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tmp := forEachNode(c, pre, post)
		if tmp != "" {
			title = tmp
		}
	}
	if post != nil {
		post(n)
	}
	return title

}
func Extract(url string) ([]string, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, "", fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, "", fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				a := link.Scheme + "://" + link.Host
				//links = append(links, link.String())
				links = append(links, a)
			}
		}

		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title := n.FirstChild.Data
			//fmt.Println(title)
			return title
		}
		return ""
	}
	title := forEachNode(doc, visitNode, nil)
	return links, title, nil
}
func crawl(url string) ([]string, string) {
	//fmt.Println(url)
	list, title, err := Extract(url)
	if err != nil {
		//log.Print(err)
	}
	return list, title
}
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func main() {
	fmt.Println("vim-go")
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	detector := chardet.NewTextDetector()
	go func() { worklist <- os.Args[1:] }()
	for i := 0; i < 30; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks, title := crawl(link)
				result, _ := detector.DetectBest([]byte(title))
				fdata := ""
				if result != nil && (result.Charset == "GB-18030" || result.Charset == "zh") {
					fdata = ConvertToString(title, "gbk", "utf-8")

					//data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(title)), simplifiedchinese.GBK.NewEncoder()))
					//fdata = string(data)
				} else if result != nil && result.Charset == "Big5" {
					//data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(title)), traditionalchinese.Big5.NewEncoder()))
					//fdata = string(data)
					fdata = ConvertToString(title, "Big5", "utf-8")
				} else if result != nil && result.Charset == "Shift_JIS" {
					fdata = ConvertToString(title, "Shift_JIS", "utf-8")
				} else if result != nil && result.Charset == "EUC-JP" {
					fdata = ConvertToString(title, "EUC-JP", "utf-8")
				} else {
					fdata = title
				}
				if result != nil {
					fmt.Printf("%s %s %s\n", link, result.Charset, strings.TrimSpace(fdata))
				} else {
					fmt.Printf("%s %s\n", link, strings.TrimSpace(fdata))
				}

				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
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
