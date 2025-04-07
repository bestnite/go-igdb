package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type GameEngineLogos struct {
	BaseEndpoint[pb.GameEngineLogo]
}

func NewGameEngineLogos(request func(URL string, dataBody any) (*resty.Response, error)) *GameEngineLogos {
	a := &GameEngineLogos{
		BaseEndpoint[pb.GameEngineLogo]{
			endpointName: EPGameEngineLogos,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameEngineLogos) Query(query string) ([]*pb.GameEngineLogo, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameEngineLogoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gameenginelogos) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gameenginelogos, nil
}
