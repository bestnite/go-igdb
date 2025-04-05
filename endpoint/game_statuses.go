package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type GameStatuses struct {
	BaseEndpoint[pb.GameStatus]
}

func NewGameStatuses(request func(URL string, dataBody any) (*resty.Response, error)) *GameStatuses {
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
