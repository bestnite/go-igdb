package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetGameStatuses(query string) ([]*pb.GameStatus, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_statuses.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameStatusResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gamestatuses) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gamestatuses, nil
}

func (g *Client) GetGameStatusByID(id uint64) (*pb.GameStatus, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameStatuses, err := g.GetGameStatuses(query)
	if err != nil {
		return nil, err
	}
	return gameStatuses[0], nil
}

func (g *Client) GetGameStatusesByIDs(ids []uint64) ([]*pb.GameStatus, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameStatuses(idStr)
}

func (g *Client) GetGameStatusesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	gameStatuses, err := g.GetGameStatuses(query)
	if err != nil {
		return 0, err
	}
	return int(gameStatuses[0].Id), nil
}
