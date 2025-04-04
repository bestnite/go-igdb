package igdb

import (
	"fmt"
	"strings"

	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGames(query string) ([]*pb.Game, error) {
	resp, err := g.Request("https://api.igdb.com/v4/games.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Games) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Games, nil
}

func (g *igdb) GetGameByID(id uint64) (*pb.Game, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	games, err := g.GetGames(query)
	if err != nil {
		return nil, err
	}
	return games[0], nil
}

func (g *igdb) GetGameByIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}
