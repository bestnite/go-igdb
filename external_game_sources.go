package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetExternalGameSources(query string) ([]*pb.ExternalGameSource, error) {
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

func (g *igdb) GetExternalGameSourceByID(id uint64) (*pb.ExternalGameSource, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	externalGameSources, err := g.GetExternalGameSources(query)
	if err != nil {
		return nil, err
	}
	return externalGameSources[0], nil
}

func (g *igdb) GetExternalGameSourcesByIDs(ids []uint64) ([]*pb.ExternalGameSource, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetExternalGameSources(idStr)
}
