package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type ReleaseDateStatuses struct {
	BaseEndpoint[pb.ReleaseDateStatus]
}

func NewReleaseDateStatuses(request RequestFunc) *ReleaseDateStatuses {
	a := &ReleaseDateStatuses{
		BaseEndpoint: BaseEndpoint[pb.ReleaseDateStatus]{
			endpointName: EPReleaseDateStatuses,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *ReleaseDateStatuses) Query(ctx context.Context, query string) ([]*pb.ReleaseDateStatus, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ReleaseDateStatusResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Releasedatestatuses, nil
}
