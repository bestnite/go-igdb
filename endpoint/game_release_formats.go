package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameReleaseFormats struct {
	BaseEndpoint[pb.GameReleaseFormat]
}

func NewGameReleaseFormats(request RequestFunc) *GameReleaseFormats {
	a := &GameReleaseFormats{
		BaseEndpoint[pb.GameReleaseFormat]{
			endpointName: EPGameReleaseFormats,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameReleaseFormats) Query(query string) ([]*pb.GameReleaseFormat, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameReleaseFormatResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gamereleaseformats) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gamereleaseformats, nil
}
