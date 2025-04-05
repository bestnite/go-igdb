package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetKeywords(query string) ([]*pb.Keyword, error) {
	resp, err := g.Request("https://api.igdb.com/v4/keywords.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.KeywordResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Keywords) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Keywords, nil
}

func (g *Client) GetKeywordByID(id uint64) (*pb.Keyword, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	keywords, err := g.GetKeywords(query)
	if err != nil {
		return nil, err
	}
	return keywords[0], nil
}

func (g *Client) GetKeywordsByIDs(ids []uint64) ([]*pb.Keyword, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetKeywords(idStr)
}

func (g *Client) GetKeywordsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	keywords, err := g.GetKeywords(query)
	if err != nil {
		return 0, err
	}
	return int(keywords[0].Id), nil
}
