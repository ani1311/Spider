package main

import "fmt"

func main() {

	// 137616
	links := NewSet()

	// curLinks := []string{"https://en.wikipedia.org/wiki/Main_Page"}
	curLinks := []string{"https://github.com/"}

	for i := 0; i < 3; i++ {
		nextCurLinks := []string{}

		for _, link := range curLinks {
			if links.Contains(link) {
				continue
			}
			links.Add(link)
			for _, newLink := range getLinksInPage(link) {
				fmt.Println(newLink)
				nextCurLinks = append(nextCurLinks, newLink)
			}
		}
		curLinks = nextCurLinks
	}
	fmt.Println(len(links.m))
}
