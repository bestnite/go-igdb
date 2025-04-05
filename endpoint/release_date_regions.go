package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type ReleaseDateRegions struct {
	BaseEndpoint[pb.ReleaseDateRegion]
}

func NewReleaseDateRegions(request func(URL string, dataBody any) (*resty.Response, error)) *ReleaseDateRegions {
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
	resp, err := a.request("https://api.igdb.com/v4/release_date_regions.pb", query)
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
