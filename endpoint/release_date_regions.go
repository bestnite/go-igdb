package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type ReleaseDateRegions struct {
	BaseEndpoint[pb.ReleaseDateRegion]
}

func NewReleaseDateRegions(request RequestFunc) *ReleaseDateRegions {
	a := &ReleaseDateRegions{
		BaseEndpoint[pb.ReleaseDateRegion]{
			endpointName: EPReleaseDateRegions,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *ReleaseDateRegions) Query(query string) ([]*pb.ReleaseDateRegion, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ReleaseDateRegionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Releasedateregions) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Releasedateregions, nil
}
