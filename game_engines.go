package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetGameEngines(query string) ([]*pb.GameEngine, error) {
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

func (g *Client) GetGameEngineByID(id uint64) (*pb.GameEngine, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameEngines, err := g.GetGameEngines(query)
	if err != nil {
		return nil, err
	}
	return gameEngines[0], nil
}

func (g *Client) GetGameEnginesByIDs(ids []uint64) ([]*pb.GameEngine, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameEngines(idStr)
}

func (g *Client) GetGameEnginesByGameID(id uint64) ([]*pb.GameEngine, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetGameEngines(query)
}

func (g *Client) GetGameEnginesByGameIDs(ids []uint64) ([]*pb.GameEngine, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameEngines(idStr)
}

func (g *Client) GetGameEnginesByLogoID(id uint64) ([]*pb.GameEngine, error) {
	query := fmt.Sprintf(`where logo = %d; fields *;`, id)
	return g.GetGameEngines(query)
}

func (g *Client) GetGameEnginesByLogoIDs(ids []uint64) ([]*pb.GameEngine, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where logo = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameEngines(idStr)
}

func (g *Client) GetGameEnginesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	gameEngines, err := g.GetGameEngines(query)
	if err != nil {
		return 0, err
	}
	return int(gameEngines[0].Id), nil
}
