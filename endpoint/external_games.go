package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type ExternalGames struct {
	BaseEndpoint[pb.ExternalGame]
}

func NewExternalGames(request func(URL string, dataBody any) (*resty.Response, error)) *ExternalGames {
	a := &ExternalGames{
		BaseEndpoint[pb.ExternalGame]{
			endpointName: EPExternalGames,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *ExternalGames) Query(query string) ([]*pb.ExternalGame, error) {
	resp, err := a.request("https://api.igdb.com/v4/external_games.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ExternalGameResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Externalgames) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Externalgames, nil
}
