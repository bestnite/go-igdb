package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetLanguageSupports(query string) ([]*pb.LanguageSupport, error) {
	resp, err := g.Request("https://api.igdb.com/v4/language_supports.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.LanguageSupportResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Languagesupports) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Languagesupports, nil
}

func (g *Client) GetLanguageSupportByID(id uint64) (*pb.LanguageSupport, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	languageSupports, err := g.GetLanguageSupports(query)
	if err != nil {
		return nil, err
	}
	return languageSupports[0], nil
}

func (g *Client) GetLanguageSupportsByIDs(ids []uint64) ([]*pb.LanguageSupport, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetLanguageSupports(idStr)
}

func (g *Client) GetLanguageSupportsByGameID(id uint64) ([]*pb.LanguageSupport, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetLanguageSupports(query)
}

func (g *Client) GetLanguageSupportsByGameIDs(ids []uint64) ([]*pb.LanguageSupport, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetLanguageSupports(idStr)
}

func (g *Client) GetLanguageSupportsByLanguageID(id uint64) ([]*pb.LanguageSupport, error) {
	query := fmt.Sprintf(`where language = %d; fields *;`, id)
	return g.GetLanguageSupports(query)
}

func (g *Client) GetLanguageSupportsByLanguageIDs(ids []uint64) ([]*pb.LanguageSupport, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where language = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetLanguageSupports(idStr)
}

func (g *Client) GetLanguageSupportsByLanguageSupportTypeID(id uint64) ([]*pb.LanguageSupport, error) {
	query := fmt.Sprintf(`where language_support_type = %d; fields *;`, id)
	return g.GetLanguageSupports(query)
}

func (g *Client) GetLanguageSupportsByLanguageSupportTypeIDs(ids []uint64) ([]*pb.LanguageSupport, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where language_support_type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetLanguageSupports(idStr)
}

func (g *Client) GetLanguageSupportsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	languageSupports, err := g.GetLanguageSupports(query)
	if err != nil {
		return 0, err
	}
	return int(languageSupports[0].Id), nil
}
