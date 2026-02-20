package feedhandler

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mmcdole/gofeed"
	// ext "github.com/mmcdole/gofeed/extensions"
)

type FeedHandler struct {
	Feeds   []Feed
	FeedMap map[string]Feed
}

func NewFeedHandler() *FeedHandler {
	fh := &FeedHandler{
		FeedMap: make(map[string]Feed),
	}
	return fh
}

func (fh *FeedHandler) Summary() string {
	result := strings.Builder{}
	for _, feed := range fh.Feeds {
		result.WriteString(feed.Summary() + "\n")
	}
	return result.String()
}

// FetchFeed fetches an RSS feed from the given URL and returns the response body
func FetchFeed(feedURL string) ([]byte, error) {
	fmt.Printf("Fetching feed from: %s\n\n", feedURL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:91.0) Gecko/20100101 Firefox/91.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching feed: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d %s", resp.StatusCode, resp.Status)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return body, nil
}

// func (fh *FeedHandler) 

func (fh *FeedHandler) AddFeed(toAdd Feed) error {
	fh.Feeds = append(fh.Feeds, toAdd)
	fh.FeedMap[toAdd.Name] = toAdd
	return nil
}

// ParseRSSGofeed parses RSS/Atom feeds using the gofeed library
// This is a production-ready parser that handles multiple feed formats
func ParseRSSGofeed(data []byte) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse feed: %w", err)
	}
	return feed, nil
}
