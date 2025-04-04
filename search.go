package igdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	pb "github/bestnite/go-igdb/proto"

	"github.com/PuerkitoBio/goquery"
	"github.com/bestnite/go-flaresolverr"
	"google.golang.org/protobuf/proto"
)

var webSearchCFCookies struct {
	cookies []*http.Cookie
	expires time.Time
}

func (g *igdb) Search(query string) ([]*pb.Search, error) {
	resp, err := g.Request("https://api.igdb.com/v4/search.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.SearchResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Searches) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Searches, nil
}

func (g *igdb) WebSearchGames(name string) ([]*pb.Game, error) {
	params := url.Values{}
	params.Add("q", name)
	params.Add("utf8", "✓")
	Url := fmt.Sprintf("%s?%s", "https://www.igdb.com/search", params.Encode())

	f, err := g.getFlaresolverr()
	if err != nil {
		return nil, fmt.Errorf("failed to get flaresolverr: %w", err)
	}

	var respBody io.Reader
	if webSearchCFCookies.cookies == nil || time.Now().After(webSearchCFCookies.expires) {
		resp, err := f.GetV1(Url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to web search: %s: %w", name, err)
		}
		webSearchCFCookies.cookies = resp.Solution.Cookies
		webSearchCFCookies.expires = time.Now().Add(3 * time.Hour)
		respBody = strings.NewReader(resp.Solution.Response)
	} else if time.Now().Before(webSearchCFCookies.expires) {
		resp, err := f.SimulateGet(Url, &flaresolverr.SimulateOptions{
			HttpCookies: webSearchCFCookies.cookies,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to search IGDB ID: %s: %w", name, err)
		}
		respBody = strings.NewReader(resp.Body)
	}

	doc, err := goquery.NewDocumentFromReader(respBody)
	if err != nil {
		return nil, fmt.Errorf("failed to parse IGDB web search response: %w", err)
	}
	selection := doc.Find("script[data-component-name='GameEntries']")
	if selection.Length() == 0 {
		return nil, fmt.Errorf("no search results found for: %s", name)
	}
	innerJson := selection.First().Text()
	data := struct {
		Games []struct {
			Id   uint64 `json:"id"`
			Name string `json:"name"`
		} `json:"games"`
	}{}
	if err := json.Unmarshal([]byte(innerJson), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal IGDB web search response: %w", err)
	}

	ids := make([]uint64, len(data.Games))
	for i, game := range data.Games {
		ids[i] = game.Id
	}

	return g.GetGameByIDs(ids)
}
