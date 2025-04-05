package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameTimeToBeats struct{ BaseEndpoint }

func (a *GameTimeToBeats) Query(query string) ([]*pb.GameTimeToBeat, error) {
	resp, err := a.request("https://api.igdb.com/v4/game_time_to_beats.pb", query)
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
