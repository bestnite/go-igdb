package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type GameVersions struct {
	BaseEndpoint[pb.GameVersion]
}

func NewGameVersions(request func(URL string, dataBody any) (*resty.Response, error)) *GameVersions {
	a := &GameVersions{
		BaseEndpoint[pb.GameVersion]{
			endpointName: EPGameVersions,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameVersions) Query(query string) ([]*pb.GameVersion, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameVersionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gameversions) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gameversions, nil
}
