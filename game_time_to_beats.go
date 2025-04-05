package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetGameTimeToBeats(query string) ([]*pb.GameTimeToBeat, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_time_to_beats.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameTimeToBeatResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gametimetobeats) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gametimetobeats, nil
}

func (g *Client) GetGameTimeToBeatByID(id uint64) (*pb.GameTimeToBeat, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameTimeToBeats, err := g.GetGameTimeToBeats(query)
	if err != nil {
		return nil, err
	}
	return gameTimeToBeats[0], nil
}

func (g *Client) GetGameTimeToBeatsByIDs(ids []uint64) ([]*pb.GameTimeToBeat, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameTimeToBeats(idStr)
}

func (g *Client) GetGameTimeToBeatsByGameID(id uint64) ([]*pb.GameTimeToBeat, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetGameTimeToBeats(query)
}

func (g *Client) GetGameTimeToBeatsByGameIDs(ids []uint64) ([]*pb.GameTimeToBeat, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameTimeToBeats(idStr)
}

func (g *Client) GetGameTimeToBeatsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	gameTimeToBeats, err := g.GetGameTimeToBeats(query)
	if err != nil {
		return 0, err
	}
	return int(gameTimeToBeats[0].Id), nil
}
