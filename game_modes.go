package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetGameModes(query string) ([]*pb.GameMode, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_modes.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameModeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gamemodes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gamemodes, nil
}

func (g *Client) GetGameModeByID(id uint64) (*pb.GameMode, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameModes, err := g.GetGameModes(query)
	if err != nil {
		return nil, err
	}
	return gameModes[0], nil
}

func (g *Client) GetGameModesByIDs(ids []uint64) ([]*pb.GameMode, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameModes(idStr)
}

func (g *Client) GetGameModesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	gameModes, err := g.GetGameModes(query)
	if err != nil {
		return 0, err
	}
	return int(gameModes[0].Id), nil
}
