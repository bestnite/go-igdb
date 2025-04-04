package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGameTimeToBeats(query string) ([]*pb.GameTimeToBeat, error) {
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

func (g *igdb) GetGameTimeToBeatByID(id uint64) (*pb.GameTimeToBeat, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameTimeToBeats, err := g.GetGameTimeToBeats(query)
	if err != nil {
		return nil, err
	}
	return gameTimeToBeats[0], nil
}

func (g *igdb) GetGameTimeToBeatsByIDs(ids []uint64) ([]*pb.GameTimeToBeat, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameTimeToBeats(idStr)
}
