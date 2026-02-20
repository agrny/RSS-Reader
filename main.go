package main

import (
	"flag"
	"fmt"

	feedhandler "RSS-Reader/feedHandler"
)

func main() {
	// Command-line flags
	fh := feedhandler.NewFeedHandler()
	feedURL := flag.String("a", "", "RSS feed URL")
	flag.String("add", "", "RSS feed URL")
	// maxItems := flag.Int("items", 5, "Maximum number of items to display")
	flag.Parse()

	if feedURL != nil {
		feed := feedhandler.NewFeed(*feedURL)
		fh.AddFeed(*feed)
	}

	fmt.Printf("SUMMARY\n%s", fh.Summary())

	// get feed title via channel.title
}
