package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type PopularityPrimitives struct {
	BaseEndpoint[pb.PopularityPrimitive]
}

func NewPopularityPrimitives(request RequestFunc) *PopularityPrimitives {
	a := &PopularityPrimitives{
		BaseEndpoint[pb.PopularityPrimitive]{
			endpointName: EPPopularityPrimitives,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PopularityPrimitives) Query(ctx context.Context, query string) ([]*pb.PopularityPrimitive, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PopularityPrimitiveResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Popularityprimitives, nil
}
