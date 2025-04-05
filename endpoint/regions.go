package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Regions struct {
	BaseEndpoint[pb.Region]
}

func NewRegions(request func(URL string, dataBody any) (*resty.Response, error)) *Regions {
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
	resp, err := a.request("https://api.igdb.com/v4/regions.pb", query)
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
