package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,		
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed to create a get request: %v", err) 
	}
	req.Header.Set("User-Agent", "gator_from_bootdev_user")

	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("there was an error sending your request, try again: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error reading xml doc: %w", err)
	}

	var rssString *RSSFeed
	err = xml.Unmarshal(body, &rssString)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshalling xml: %w", err)
	}

	cleanedTitleChan := html.UnescapeString(rssString.Channel.Title)
	cleanedDesChan := html.UnescapeString(rssString.Channel.Description)
	rssString.Channel.Title = cleanedTitleChan
	rssString.Channel.Description = cleanedDesChan
	for idx, feed := range rssString.Channel.Item {
		cleanedTitleFeed := html.UnescapeString(feed.Title)
		cleanedDesFeed := html.UnescapeString(feed.Description)
		rssString.Channel.Item[idx].Title = cleanedTitleFeed
		rssString.Channel.Item[idx].Description = cleanedDesFeed
	}
	return rssString, nil
}

	
// Feed was generated 2025-10-10 15:10:07 by https://xml-to-go.github.io/ in Ukraine.
/*type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Text    string   `xml:",chardata"`
	Title   string   `xml:"title"`
	Link    []struct {
		Text string `xml:",chardata"`
		Href string `xml:"href,attr"`
		Rel  string `xml:"rel,attr"`
	} `xml:"link"`
	ID      string `xml:"id"`
	Updated string `xml:"updated"`
	Entry   []struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
		} `xml:"link"`
		Published string `xml:"published"`
		Updated   string `xml:"updated"`
		Author    struct {
			Text string `xml:",chardata"`
			Name string `xml:"name"`
		} `xml:"author"`
		ID      string `xml:"id"`
		Summary struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"summary"`
		Content struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"content"`
		Category []struct {
			Text string `xml:",chardata"`
			Term string `xml:"term,attr"`
		} `xml:"category"`
	} `xml:"entry"`
} 
*/
