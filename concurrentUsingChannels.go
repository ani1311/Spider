package main

import (
	"time"
)

const NO_OF_WORKERS = 100

func worker(tpq *urlQueue, pq *urlQueue) {
	for {
		valid, link := tpq.pop()
		if !valid {
			continue
		}

		for _, newLink := range getLinksInPage(link) {
			for !pq.push(newLink) {
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}

func master(url string) {
	toProcessesQ := NewUrlQueue()
	processedQ := NewUrlQueue()
	ws := NewWebIndexSet()

	toProcessesQ.push(url)

	for i := 1; i <= NO_OF_WORKERS; i++ {
		go worker(toProcessesQ, processedQ)
	}

	for {
		valid, link := processedQ.pop()
		if !valid || ws.Contains(link) {
			continue
		}
		ws.Add(link)
		toProcessesQ.push(link)
		if len(ws.m) > 100_000 {
			break
		}
	}

}
