package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetExternalGameSources(query string) ([]*pb.ExternalGameSource, error) {
	resp, err := g.Request("https://api.igdb.com/v4/external_game_sources.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ExternalGameSourceResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Externalgamesources) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Externalgamesources, nil
}

func (g *Client) GetExternalGameSourceByID(id uint64) (*pb.ExternalGameSource, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	externalGameSources, err := g.GetExternalGameSources(query)
	if err != nil {
		return nil, err
	}
	return externalGameSources[0], nil
}

func (g *Client) GetExternalGameSourcesByIDs(ids []uint64) ([]*pb.ExternalGameSource, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetExternalGameSources(idStr)
}

func (g *Client) GetExternalGameSourcesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	externalGameSources, err := g.GetExternalGameSources(query)
	if err != nil {
		return 0, err
	}
	return int(externalGameSources[0].Id), nil
}
