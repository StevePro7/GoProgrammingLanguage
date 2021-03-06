package main

import (
	"fmt"
	"log"

	"learning/links"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	// Crawl the web breadth-first
	// starting from the command line arguments
	//breadthFirst(crawl, os.Args[1:])

	//worklist := []string {"http://www.google.com"}
	var worklist []string
	worklist = append(worklist, "http://www.google.com")
	breadthFirst(crawl, worklist)
}
