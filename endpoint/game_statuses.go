package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameStatuses struct{ BaseEndpoint }

func (a *GameStatuses) Query(query string) ([]*pb.GameStatus, error) {
	resp, err := a.request("https://api.igdb.com/v4/game_statuses.pb", query)
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
