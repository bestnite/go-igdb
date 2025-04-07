package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type GameEngines struct {
	BaseEndpoint[pb.GameEngine]
}

func NewGameEngines(request func(URL string, dataBody any) (*resty.Response, error)) *GameEngines {
	a := &GameEngines{
		BaseEndpoint[pb.GameEngine]{
			endpointName: EPGameEngines,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameEngines) Query(query string) ([]*pb.GameEngine, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameEngineResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gameengines) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gameengines, nil
}
