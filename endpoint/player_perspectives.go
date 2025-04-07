package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type PlayerPerspectives struct {
	BaseEndpoint[pb.PlayerPerspective]
}

func NewPlayerPerspectives(request func(URL string, dataBody any) (*resty.Response, error)) *PlayerPerspectives {
	a := &PlayerPerspectives{
		BaseEndpoint[pb.PlayerPerspective]{
			endpointName: EPPlayerPerspectives,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlayerPerspectives) Query(query string) ([]*pb.PlayerPerspective, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
