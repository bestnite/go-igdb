package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Regions struct {
	BaseEndpoint[pb.Region]
}

func NewRegions(request RequestFunc) *Regions {
	a := &Regions{
		BaseEndpoint[pb.Region]{
			endpointName: EPRegions,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Regions) Query(query string) ([]*pb.Region, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.RegionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Regions) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Regions, nil
}
