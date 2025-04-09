package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameStatuses struct {
	BaseEndpoint[pb.GameStatus]
}

func NewGameStatuses(request RequestFunc) *GameStatuses {
	a := &GameStatuses{
		BaseEndpoint[pb.GameStatus]{
			endpointName: EPGameStatuses,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameStatuses) Query(query string) ([]*pb.GameStatus, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
