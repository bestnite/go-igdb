package igdb

import (
	"fmt"
	"strings"

	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGameEngines(query string) ([]*pb.GameEngine, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_engines.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameEngineResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gameengines) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gameengines, nil
}

func (g *igdb) GetGameEngineByID(id uint64) (*pb.GameEngine, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameEngines, err := g.GetGameEngines(query)
	if err != nil {
		return nil, err
	}
	return gameEngines[0], nil
}

func (g *igdb) GetGameEnginesByIDs(ids []uint64) ([]*pb.GameEngine, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameEngines(idStr)
}

func (g *igdb) GetGameEnginesByGameID(id uint64) ([]*pb.GameEngine, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetGameEngines(query)
}

func (g *igdb) GetGameEnginesByGameIDs(ids []uint64) ([]*pb.GameEngine, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameEngines(idStr)
}

func (g *igdb) GetGameEnginesByLogoID(id uint64) ([]*pb.GameEngine, error) {
	query := fmt.Sprintf(`where logo = %d; fields *;`, id)
	return g.GetGameEngines(query)
}

func (g *igdb) GetGameEnginesByLogoIDs(ids []uint64) ([]*pb.GameEngine, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where logo = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameEngines(idStr)
}
