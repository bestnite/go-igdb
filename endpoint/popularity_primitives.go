package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type PopularityPrimitives struct {
	BaseEndpoint[pb.PopularityPrimitive]
}

func NewPopularityPrimitives(request func(URL string, dataBody any) (*resty.Response, error)) *PopularityPrimitives {
	a := &PopularityPrimitives{
		BaseEndpoint[pb.PopularityPrimitive]{
			endpointName: EPPopularityPrimitives,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PopularityPrimitives) Query(query string) ([]*pb.PopularityPrimitive, error) {
	resp, err := a.request("https://api.igdb.com/v4/popularity_primitives.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PopularityPrimitiveResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Popularityprimitives) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Popularityprimitives, nil
}
