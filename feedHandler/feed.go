package feedhandler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Name       string
	URL        string
	ParsedFeed *gofeed.Feed
}

func NewFeed(url string) *Feed {
	f := &Feed{
		URL: url,
	}
	f.fetchXML()
	return f
}

func (f *Feed) SetURL(url string) {
	f.URL = url
}

func (f *Feed) SetName(name string) {
	f.Name = name
}

func (f *Feed) SetParsedFeed(parsed *gofeed.Feed) {
	f.Name = parsed.Title
	f.ParsedFeed = parsed
}

func (f *Feed) Summary() string {
	itemCount := 0
	title := "Not parsed"
	if f.ParsedFeed != nil {
		itemCount = len(f.ParsedFeed.Items)
		title = f.Name
	}
	return fmt.Sprintf("Feed{\n\tName: %s\n\tURL: %s\n\tItems: %d\n}", title, f.URL, itemCount)
}

func (f *Feed) XMLString() string {
	return f.ParsedFeed.String()
}

func (f *Feed) fetchXML() (*gofeed.Feed, error) {
	feedURL := f.URL
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

	result, err := parseFeed(body)
	if err != nil {
		return nil, err
	}
	f.SetParsedFeed(result)

	return result, nil
}

func (f *Feed) Fetch(url string) (*gofeed.Feed, error) {
	feedURL := url
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

	result, err := parseFeed(body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// parseFeed parses RSS/Atom feeds using gofeed
func parseFeed(data []byte) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse feed: %w", err)
	}
	return feed, nil
}
