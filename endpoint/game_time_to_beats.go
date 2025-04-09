package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameTimeToBeats struct {
	BaseEndpoint[pb.GameTimeToBeat]
}

func NewGameTimeToBeats(request RequestFunc) *GameTimeToBeats {
	a := &GameTimeToBeats{
		BaseEndpoint[pb.GameTimeToBeat]{
			endpointName: EPGameTimeToBeats,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameTimeToBeats) Query(query string) ([]*pb.GameTimeToBeat, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
