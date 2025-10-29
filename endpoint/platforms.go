package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Platforms struct {
	BaseEndpoint[pb.Platform]
}

func NewPlatforms(request RequestFunc) *Platforms {
	a := &Platforms{
		BaseEndpoint[pb.Platform]{
			endpointName: EPPlatforms,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Platforms) Query(ctx context.Context, query string) ([]*pb.Platform, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Platforms, nil
}
