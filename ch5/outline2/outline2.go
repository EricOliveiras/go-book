package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: programa <URL>")
		return
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição HTTP:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Erro ao analisar o HTML:", err)
		return
	}

	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, pos func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, pos)
	}

	if pos != nil {
		pos(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		depth--
	}
}
