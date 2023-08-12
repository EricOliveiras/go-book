package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, link := range os.Args[1:] {
		words, images, err := CountWordsAndImages(link)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Link: %s\nwords: %d, images: %d\n-----------------------\n", link, words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childWords, childImages := countWordsAndImages(c)
		words += childWords
		images += childImages
	}
	return
}