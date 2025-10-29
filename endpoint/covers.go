package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Covers struct {
	BaseEndpoint[pb.Cover]
}

func NewCovers(request RequestFunc) *Covers {
	a := &Covers{
		BaseEndpoint[pb.Cover]{
			endpointName: EPCovers,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Covers) Query(ctx context.Context, query string) ([]*pb.Cover, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CoverResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Covers, nil
}
