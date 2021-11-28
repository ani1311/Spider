package main

func sequentialCrawl(url string) {

	wis := NewWebIndexSet()
	urls := []string{url}

	for {
		newUrls := []string{}
		for _, link := range urls {
			if wis.Contains(link) {
				continue
			}
			wis.Add(link)
			newUrls = append(newUrls, getLinksInPage(link)...)
		}
		urls = newUrls
	}
}
