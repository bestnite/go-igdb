package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGameTypes(query string) ([]*pb.GameType, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gametypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gametypes, nil
}

func (g *igdb) GetGameTypeByID(id uint64) (*pb.GameType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameTypes, err := g.GetGameTypes(query)
	if err != nil {
		return nil, err
	}
	return gameTypes[0], nil
}

func (g *igdb) GetGameTypesByIDs(ids []uint64) ([]*pb.GameType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameTypes(idStr)
}

func (g *igdb) GetGameTypesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	gameTypes, err := g.GetGameTypes(query)
	if err != nil {
		return 0, err
	}
	return int(gameTypes[0].Id), nil
}
