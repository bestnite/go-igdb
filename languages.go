package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetLanguages(query string) ([]*pb.Language, error) {
	resp, err := g.Request("https://api.igdb.com/v4/languages.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.LanguageResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Languages) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Languages, nil
}

func (g *Client) GetLanguageByID(id uint64) (*pb.Language, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	languages, err := g.GetLanguages(query)
	if err != nil {
		return nil, err
	}
	return languages[0], nil
}

func (g *Client) GetLanguagesByIDs(ids []uint64) ([]*pb.Language, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetLanguages(idStr)
}

func (g *Client) GetLanguagesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	languages, err := g.GetLanguages(query)
	if err != nil {
		return 0, err
	}
	return int(languages[0].Id), nil
}
