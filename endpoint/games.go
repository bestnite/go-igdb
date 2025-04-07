package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Games struct {
	BaseEndpoint[pb.Game]
}

func NewGames(request func(URL string, dataBody any) (*resty.Response, error)) *Games {
	a := &Games{
		BaseEndpoint[pb.Game]{
			endpointName: EPGames,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Games) Query(query string) ([]*pb.Game, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Games) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Games, nil
}
