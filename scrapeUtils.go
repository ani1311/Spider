package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func processLink(domain string, url string) (bool, string) {
	queryParamIndex := strings.IndexByte(url, '?')
	var finalUrl string
	if queryParamIndex != -1 {
		finalUrl = url[:queryParamIndex]
	} else {
		finalUrl = url
	}
	if len(finalUrl) == 0 || (finalUrl[0] != '/' && finalUrl[0] != 'h') {
		return false, ""
	} else if finalUrl[0] == '/' {
		finalUrl = domain + finalUrl
	}
	return true, finalUrl
}

func getLinksFromHref(domain string, attributes []html.Attribute) []string {
	links := []string{}

	for _, atr := range attributes {
		if atr.Key == "href" {
			valid, url := processLink(domain, atr.Val)
			if valid {
				links = append(links, url)
			}
		}
	}

	return links
}

func getLinksInPage(linkUrl string) []string {
	resp, err := http.Get(linkUrl)

	if err != nil {
		log.Println(err)
		return []string{}
	}

	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

	urlParse, err := url.Parse(linkUrl)
	if err != nil {
		panic(err)
	}

	domain := urlParse.Scheme + "://" + urlParse.Host

	links := []string{}

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return links
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				links = append(links, getLinksFromHref(domain, t.Attr)...)
			}
		}
	}
	return links
}
