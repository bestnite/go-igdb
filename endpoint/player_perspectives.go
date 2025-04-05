package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type PlayerPerspectives struct{ BaseEndpoint }

func (a *PlayerPerspectives) Query(query string) ([]*pb.PlayerPerspective, error) {
	resp, err := a.request("https://api.igdb.com/v4/player_perspectives.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlayerPerspectiveResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Playerperspectives) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Playerperspectives, nil
}
