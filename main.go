package main

import (
	"flag"
	"fmt"

	feedhandler "RSS-Reader/feedHandler"
)

func main() {
	// Command-line flags
	defaultFeedURL := "https://www.wired.com/feed/rss"
	feedURL := flag.String("u", "", "RSS feed URL")
	flag.String("url", defaultFeedURL, "RSS feed URL")
	// maxItems := flag.Int("items", 5, "Maximum number of items to display")
	flag.Parse()

	// Fetch RSS feed
	feed := feedhandler.NewFeed(*feedURL)

	fmt.Printf("SUMMARY:\n\n%s", feed.Summary())

	// get feed title via channel.title
}
