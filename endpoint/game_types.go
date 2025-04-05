package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type GameTypes struct {
	BaseEndpoint[pb.GameType]
}

func NewGameTypes(request func(URL string, dataBody any) (*resty.Response, error)) *GameTypes {
	a := &GameTypes{
		BaseEndpoint[pb.GameType]{
			endpointName: EPGameTypes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameTypes) Query(query string) ([]*pb.GameType, error) {
	resp, err := a.request("https://api.igdb.com/v4/game_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gametypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gametypes, nil
}
